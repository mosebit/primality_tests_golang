package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// проверка полученного числа больше ли оно 5 (по требованию алгоритма теста Ферма)
func check_more_then_five(number_str string) bool {
	five_num := big.NewInt(4)
	main_num := new(big.Int)
	main_num.SetString(number_str, 10)

	if main_num.Cmp(five_num) > 0 {
		return true
	} else {
		return false
	}
}

func ferma(number_str string, count int) {
	fmt.Println("\n--------------------------")
	fmt.Printf("Тест Ферма для числа <%s>\n", number_str)

	if check_more_then_five(number_str) == false {
		fmt.Println("Число должно быть > 5")
		return
	}

	n := new(big.Int)
	n.SetString(number_str, 10)

	// a := new(big.Int)
	n_sub_two := big.NewInt(0).Sub(n, big.NewInt(2))
	bases := []string{}
	for i := 0; i != count; i++ {
		rand := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		a := big.NewInt(0).Rand(rand, n_sub_two)
		r := big.NewInt(0).Exp(a, big.NewInt(0).Sub(n, big.NewInt(1)), n)

		// fmt.Printf("r:%s i:%d\n", r.String(), i)
		if r.Cmp(big.NewInt(1)) == 0 {
			// fmt.Printf("Число <%s>, вероятно, простое!", number_str)
			bases = append(bases, a.String())
			time.Sleep(2 * time.Millisecond)
		} else {
			fmt.Printf("Число <%s> составное, основание: r!=1\n", number_str)
			return
		}
	}

	fmt.Printf("Число <%s>, вероятно, простое. Основания:\n", number_str)
	for i := 0; i != 5; {
		fmt.Printf("%s\n", bases[i])
		i++
	}
}

func solovay_strassen(number_str string, checks_count int) {
	fmt.Println("\n--------------------------")
	fmt.Printf("Тест Соловэя-Штрассена для числа <%s>\n", number_str)

	if check_more_then_five(number_str) == false {
		fmt.Println("Число должно быть > 5")
		return
	}
	n := new(big.Int)
	n.SetString(number_str, 10)

	n_sub_two := big.NewInt(0).Sub(n, big.NewInt(2))
	bases := []string{}
	for i := 0; i != checks_count; i++ {
		rand := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		a := big.NewInt(0).Rand(rand, n_sub_two)
		r := big.NewInt(0).Exp(a, big.NewInt(0).Div(big.NewInt(0).Sub(n, big.NewInt(1)), big.NewInt(2)), n)

		// fmt.Printf("r:%s i:%d\n", r.String(), i)
		if r.Cmp(big.NewInt(1)) != 0 && r.Cmp(big.NewInt(0).Sub(n, big.NewInt(1))) != 0 {
			// bases = append(bases, a.String())
			// time.Sleep(1 * time.Millisecond)
			fmt.Printf("Число составное, основание: %s\nПричина: r!=1 и r!=n-1\n", a.String())
			return
		}

		s := big.Jacobi(a, n)

		if r.Cmp(big.NewInt(0).Mod(big.NewInt(int64(s)), n)) != 0 {
			fmt.Printf("Число составное, основание: %s\nПричина: r!=s(mod n)\n", a.String())
			return
		} else {
			bases = append(bases, a.String())
		}
		time.Sleep(2 * time.Millisecond)
	}

	fmt.Printf("Число <%s>, вероятно, простое. Основания:\n", number_str)
	for i := 0; i != 5 && i != len(bases); i++ {
		fmt.Printf("%s\n", bases[i])
	}
}

func rabin_miller(number_str string, count int) {
	fmt.Println("\n--------------------------")
	fmt.Printf("Тест Рабина-Миллера для числа <%s>\n", number_str)

	num_ONE := big.NewInt(1)
	num_TWO := big.NewInt(2)
	n := new(big.Int)
	tmp_big := new(big.Int)

	bases := []string{}
	for k := 0; k != count; k++ {
		n.SetString(number_str, 10)
		n.Sub(n, big.NewInt(1))
		s := new(big.Int)
		r := new(big.Int)
		for {
			if r.Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
				break
			}
			n.Div(n, num_TWO)
			s.Add(s, num_ONE)
		}
		r = n
		n.SetString(number_str, 10)
		n_sub_two := big.NewInt(0).Sub(n, num_TWO)
		tmprand := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		a := big.NewInt(0).Rand(tmprand, n_sub_two)
		for a.Cmp(num_ONE) < 0 {
			time.Sleep(2 * time.Millisecond)
			tmprand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
			a = big.NewInt(0).Rand(tmprand, n_sub_two)
		}
		n.SetString(number_str, 10)
		y := big.NewInt(0).Exp(a, r, n)
		j := new(big.Int)
		for_flag := 0
		if y.Cmp(num_ONE) != 0 && y.Cmp(big.NewInt(0).Sub(n, num_ONE)) != 0 {
			// TODO исправить ошибку где-то здесь
			j = big.NewInt(1)
			for j.Cmp(tmp_big.Sub(s, num_ONE)) < 1 && y.Cmp(big.NewInt(0).Sub(n, num_ONE)) != 0 {
				for_flag = 1
				y.Exp(y, num_TWO, n)
				if y.Cmp(num_ONE) == 0 {
					// fmt.Printf("Число <%s> составное\nПричина: y==1 после y=y^2(mod n)\n", number_str)
					return
				}
				j.Add(j, num_ONE)
			}
			if for_flag == 1 && y.Cmp(big.NewInt(0).Sub(n, num_ONE)) != 0 {
				fmt.Printf("Число <%s> Составное\nПричина: y!=n-1\n", number_str)
				return
			}
			j.Add(j, num_ONE)
		}

		bases = append(bases, a.String())
		time.Sleep(2 * time.Millisecond)

	}

	fmt.Printf("Число <%s>, вероятно, простое. Основания:\n", number_str)
	for i := 0; i != 5 && i != len(bases); i++ {
		fmt.Printf("%s\n", bases[i])
	}
}

// func carmichael_dop_func(num1 *big.Int, prost1 []int) {
func carmichael_dop_func(num_str string, prost1 []int) {
	num1 := big.NewInt(0)
	num1.SetString(num_str, 10)
	n_sub_1 := big.NewInt(0)
	tmp_big := big.NewInt(0)
	ZERO := big.NewInt(0)
	ONE := big.NewInt(1)
	n_sub_1.Sub(num1, ONE)
	for _, value := range prost1 {
		tmp_big.Mod(n_sub_1, big.NewInt(int64(value-1)))
		if tmp_big.Cmp(ZERO) != 0 {
			fmt.Printf("Число <%s> НЕ является Кармайкловским по второму признаку!\n", num1.String())
			fmt.Printf("Так как <%s> не делится на <%d>\n", n_sub_1.String(), value-1)
			return
		}
	}
	fmt.Printf("Число <%s> является Кармайкловским по второму признаку!\n", num1.String())
	return
}

func carmichael_check() {
	// для проверки чисел по второму признаку, являются ли они Карлмайкловскими:
	// prost1 := []int{11, 13, 17, 19, 29, 31, 37, 41, 43, 61, 71, 73, 97, 101, 109, 113, 151, 181, 193, 641}
	// prost2 := []int{13, 17, 19, 23, 29, 31, 37, 41, 43, 61, 67, 71, 73, 109, 113, 127, 151, 281, 353}
	// prost3 := []int{13, 17, 19, 23, 29, 31, 37, 41, 43, 61, 67, 71, 73, 97, 127, 199, 281, 397}
	// carmichael_dop_func("349407515342287435050603204719587201", prost1)
	// carmichael_dop_func("2810864562635368426005268142616001", prost2)
	// carmichael_dop_func("32809426840359564991177172754241", prost3)

	// тесты с числами Кармайкла:
	num1 := new(big.Int)
	num2 := new(big.Int)
	num3 := new(big.Int)
	num1.SetString("349407515342287435050603204719587201", 10)
	num2.SetString("2810864562635368426005268142616001", 10)
	num3.SetString("32809426840359564991177172754241", 10)

	// ferma(num1.String(), 100)
	// ferma(num2.String(), 100)
	// ferma(num3.String(), 100)

	// solovay_strassen(num1.String(), 100)
	// solovay_strassen(num2.String(), 100)
	// solovay_strassen(num3.String(), 100)

	rabin_miller(num1.String(), 100)
	rabin_miller(num2.String(), 100)
	rabin_miller(num3.String(), 100)
}

func main() {
	// carmichael_check()

	num1 := new(big.Int)
	num2 := new(big.Int)
	num3 := new(big.Int)
	num4 := new(big.Int)
	num1.SetString("30384433192474324819", 10)
	num2.SetString("6428091931117869454641705454180338513613", 10)
	num3.SetString("982305677759838141204881840427402188407", 10)
	num4.SetString("58613636122607904912590478499255127058383896240896572417646117868651898721679891", 10)

	// // ferma(num1.String(), 100)
	// // ferma(num2.String(), 100)
	// // ferma(num3.String(), 100)
	// // ferma(num4.String(), 100)

	// // solovay_strassen(num1.String(), 100)
	// // solovay_strassen(num2.String(), 100)
	// // solovay_strassen(num3.String(), 100)
	// // solovay_strassen(num4.String(), 100)

	// rabin_miller(num1.String(), 100)
	// rabin_miller(num2.String(), 100)
	// rabin_miller(num3.String(), 100)
	// rabin_miller(num4.String(), 100)

	// // проверка встроенной функцией
	// if num1.ProbablyPrime(1000) == true {
	// 	fmt.Printf("<%s> NOT prime!\n", num1.String())
	// } else {
	// 	fmt.Printf("<%s> NOT prime!\n", num1.String())
	// }
	// if num2.ProbablyPrime(1000) == true {
	// 	fmt.Printf("<%s> NOT prime!\n", num2.String())
	// } else {
	// 	fmt.Printf("<%s> NOT prime!\n", num2.String())
	// }
	// if num3.ProbablyPrime(1000) == true {
	// 	fmt.Printf("<%s> NOT prime!\n", num3.String())
	// } else {
	// 	fmt.Printf("<%s> NOT prime!\n", num3.String())
	// }
	// if num4.ProbablyPrime(1000) == true {
	// 	fmt.Printf("<%s> NOT prime!\n", num4.String())
	// } else {
	// 	fmt.Printf("<%s> NOT prime!\n", num4.String())
	// }
}
