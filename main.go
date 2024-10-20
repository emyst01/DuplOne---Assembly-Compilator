package main

import (
	"com/packages/ram"
	"com/packages/stack"
	"com/packages/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stackmem = stack.Stack{}
var mem = ram.NewRAM()
var A uint8
var B uint8
var Z bool

func main() {
	cn, err := os.ReadFile("index.as")
	if err != nil {
		fmt.Println(err)
	}
	c := string(cn)
	splArr := strings.Split(c, "\n")

	run(splArr)

	fmt.Println("END")
	fmt.Scanln()
}

func run(instarr []string) {
	ind := 0
	for ind < len(instarr) {
		line := utils.CleanString(instarr[ind])
		beg := strings.Split(line, " ")
		if beg[0] == "mov" {
			switch beg[1] {
			case "a":
				if utils.IsInt(beg[2]) {
					r, _ := strconv.ParseUint(beg[2], 10, 8)
					A = uint8(r)
				} else if beg[2] == "b" {
					A = B
				} else if utils.IsValidNumberFormat(beg[2]) {
					addrStr := strings.TrimSuffix(strings.TrimPrefix(beg[2], "<"), ">")
					addr, _ := strconv.ParseUint(addrStr, 10, 16)
					A = mem.Read(uint16(addr))
				} else {
					fmt.Println("Error at line", ind)
					break
				}
			case "b":
				if utils.IsInt(beg[2]) {
					r, _ := strconv.ParseUint(beg[2], 10, 8)
					B = uint8(r)
				} else if beg[2] == "a" {
					B = A
				} else if utils.IsValidNumberFormat(beg[2]) {
					addrStr := strings.TrimSuffix(strings.TrimPrefix(beg[2], "<"), ">")
					addr, _ := strconv.ParseUint(addrStr, 10, 16)
					B = mem.Read(uint16(addr))
				} else {
					fmt.Println("Error at line", ind)
					break
				}
			default:
				if utils.IsValidNumberFormat(beg[1]) {
					if beg[2] == "a" {
						addrStr := strings.TrimSuffix(strings.TrimPrefix(beg[2], "<"), ">")
						addr, _ := strconv.ParseUint(addrStr, 10, 16)
						mem.Write(uint16(addr), A)
					} else if beg[2] == "b" {
						addrStr := strings.TrimSuffix(strings.TrimPrefix(beg[2], "<"), ">")
						addr, _ := strconv.ParseUint(addrStr, 10, 16)
						mem.Write(uint16(addr), B)
					} else {
						fmt.Println("Error at line", ind)
						break
					}
				} else {
					fmt.Println("Error at line", ind)
					break
				}
			}
		} else if beg[0] == "add" {
			if beg[1] == "a" && beg[2] == "b" {
				t := A + B
				A = t
				if t == 0 {
					Z = true
				} else {
					Z = false
				}
			} else if beg[1] == "b" && beg[2] == "a" {
				t := A + B
				B = t
				if t == 0 {
					Z = true
				} else {
					Z = false
				}
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		} else if beg[0] == "sub" {
			if beg[1] == "a" && beg[2] == "b" || beg[1] == "" {
				t := A - B
				A = t
				if t == 0 {
					Z = true
				} else {
					Z = false
				}
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		} else if beg[0] == "inc" {
			switch beg[1] {
			case "a":
				t := A + 1
				A = t
				if t == 0 {
					Z = true
				} else {
					Z = false
				}
			case "b":
				t := B + 1
				B = t
				if t == 0 {
					Z = true
				} else {
					Z = false
				}
			default:
				fmt.Println("Error at line", ind)
				return
			}
		} else if beg[0] == "dec" {
			switch beg[1] {
			case "a":
				t := A - 1
				A = t
				if t == 0 {
					Z = true
				} else {
					Z = false
				}
			case "b":
				t := B - 1
				B = t
				if t == 0 {
					Z = true
				} else {
					Z = false
				}
			default:
				fmt.Println("Error at line", ind)
				return
			}
		} else if beg[0] == "nop" {
			fmt.Println("...")
		} else if beg[0] == "halt" {
			fmt.Println("halt")
			break
		} else if beg[0] == "in" {
			var input string
			fmt.Scanln(&input)
			if utils.IsUint8(input) {
				r, _ := strconv.ParseUint(input, 10, 8)
				A = uint8(r)
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		} else if beg[0] == "out" {
			fmt.Println(A)
		} else if beg[0] == "push" {
			if beg[1] == "a" {
				stackmem.Push(A)
			} else if beg[1] == "b" {
				stackmem.Push(B)
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		} else if beg[0] == "pop" {
			if beg[1] == "a" {
				A = stackmem.Pop()
			} else if beg[1] == "b" {
				B = stackmem.Pop()
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		} else if beg[0] == "jmp" {
			if utils.IsValidNumberFormat(beg[1]) {
				r, _ := strconv.ParseUint(beg[1], 10, 16)
				j := int(mem.Read(uint16(r)))
				ind = j
				continue
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		} else if beg[0] == "jz" {
			if utils.IsValidNumberFormat(beg[1]) {
				if Z {
					r, _ := strconv.ParseUint(beg[1], 10, 16)
					j := int(mem.Read(uint16(r)))
					ind = j
					continue
				}
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		} else if beg[0] == "jnz" {
			if utils.IsValidNumberFormat(beg[1]) {
				if !Z {
					r, _ := strconv.ParseUint(beg[1], 10, 16)
					j := int(mem.Read(uint16(r)))
					ind = j
					continue
				}
			} else {
				fmt.Println("Error at line", ind)
				break
			}
		}
		ind++
	}
}
