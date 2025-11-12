package main

import "fmt"

func testSampleAddMethod() {
	fmt.Println("Test sample Add method:")
	// []int{1,1} []int{2,2} // 5,12
	s1 := sample{
		content:[]int{1,1},
	}
	s2 := sample{
		content:[]int{2,2},
	}
	test := s1.add(s2)
	fmt.Println("test expected: 3,3\tresult:",test)

	// []int{2,5} []int{3,7} // 5,12
	s1 = sample{
		content:[]int{2,5},
	}
	s2 = sample{
		content:[]int{3,7},
	}
	test = s1.add(s2)
	fmt.Println("test expected: 5,12\tresult:",test)

	// []int{-2,5} []int{10,-1} // 8,4
	s1 = sample{
		content:[]int{-2,5},
	}
	s2 = sample{
		content:[]int{10,-1},
	}
	test = s1.add(s2)
	fmt.Println("test expected: 8,4\tresult:",test)

	// []int{-1,-2} []int{-3,-4} // -4,-6
	s1 = sample{
		content:[]int{-1,-2},
	}
	s2 = sample{
		content:[]int{-3,-4},
	}
	test = s1.add(s2)
	fmt.Println("test expected: -4,-6\tresult:",test)
	fmt.Println("----------------------")
}

func testSampleMulMethod() {
	fmt.Println("Test sample Mul method:")
	// []int{1,1} []int{2,2} // 0,4
	s1 := sample{
		content:[]int{1,1},
	}
	s2 := sample{
		content:[]int{2,2},
	}
	test := s1.mul(s2)
	fmt.Println("test expected: 0,4\tresult:",test)

	// []int{2,5} []int{3,7} // -29,29
	s1 = sample{
		content:[]int{2,5},
	}
	s2 = sample{
		content:[]int{3,7},
	}
	test = s1.mul(s2)
	fmt.Println("test expected: -29,29\tresult:",test)

	// []int{-2,5} []int{10,-1} // -15,52
	s1 = sample{
		content:[]int{-2,5},
	}
	s2 = sample{
		content:[]int{10,-1},
	}
	test = s1.mul(s2)
	fmt.Println("test expected: -15,52\tresult:",test)

	// []int{-1,-2} []int{-3,-4} // -5,10
	s1 = sample{
		content:[]int{-1,-2},
	}
	s2 = sample{
		content:[]int{-3,-4},
	}
	test = s1.mul(s2)
	fmt.Println("test expected: -5,10\tresult:",test)

	fmt.Println("----------------------")
}

func testSampleDivMethod() {
	fmt.Println("Test sample Mul method:")
	// []int{10,12} []int{2,2} // 5,6
	s1 := sample{
		content:[]int{10,12},
	}
	s2 := sample{
		content:[]int{2,2},
	}
	test := s1.div(s2)
	fmt.Println("test expected: 5,6\tresult:",test)
	
	// []int{11,12} []int{3,5} // 3,2
	s1 = sample{
		content:[]int{11,12},
	}
	s2 = sample{
		content:[]int{3,5},
	}
	test = s1.div(s2)
	fmt.Println("test expected: 3,2\tresult:",test)

	// []int{-10,-12} []int{2,2} // -5,-6
	s1 = sample{
		content:[]int{-10,-12},
	}
	s2 = sample{
		content:[]int{2,2},
	}
	test = s1.div(s2)
	fmt.Println("test expected: -5,-6\tresult:",test)

	// []int{-11,-12} []int{3,5} // -3,-2
	s1 = sample{
		content:[]int{-11,-12},
	}
	s2 = sample{
		content:[]int{3,5},
	}
	test = s1.div(s2)
	fmt.Println("test expected: -3,-2\tresult:",test)

	fmt.Println("----------------------")
}
