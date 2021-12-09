package main

import "fmt"

func main() {
	fmt.Println("体脂率计算器")
	var calculateNum int32 = 0
	var averageFatRate float64
	var fatRateSum float64

	var suggestLevel1 string = "偏瘦"
	var suggestLevel2 string = "标准"
	var suggestLevel3 string = "偏重"
	var suggestLevel4 string = "肥胖"
	var suggestLevel5 string = "严重肥胖"

	for personNum := 1; personNum < 10; personNum++ {
		// 姓名
		var name string
		fmt.Print("请输入姓名(输入 0 退出): ")
		fmt.Scanln(&name)
		if name == "0" {
			break
		}

		// 性别
		var gender string
		var genderType int
		fmt.Print("请输入性别(男 b boy;  女 g girl): ")
		fmt.Scanln(&gender)
		if gender == "b" || gender == "boy" {
			genderType = 1
		} else if gender == "g" || gender == "girl" {
			genderType = 0
		} else {
			genderType = 1
		}

		// 身高
		var stature float64
		fmt.Print("请输入身高(m):")
		fmt.Scanln(&stature)

		// 体重
		var weight float64
		fmt.Print("请输入体重(kg):")
		fmt.Scanln(&weight)

		// 年龄
		var age int32
		fmt.Print("请输入年龄:")
		fmt.Scanln(&age)

		// BMI = 体重(kg) / (身高(m) * 身高(m))
		var bmi float64 = weight / (stature * stature)
		// 体脂率 = 1.2 * BMI + 0.23 * 年龄 - 5.4 - 10.8 * 性别 (男-1 女-0)
		var fatRate = 1.2 * bmi + 0.23 * float64(age) - 5.4 - 10.8 * float64(genderType)
		fatRateSum += fatRate

		var suggest string
		if genderType == 1 {
			// 男 体脂率建议
			if age >= 18 && age <=39 {
				if fatRate >= 0 && fatRate <= 10 {
					suggest = suggestLevel1
				} else if fatRate > 10 && fatRate <= 16 {
					suggest = suggestLevel2
				} else if fatRate > 16 && fatRate <= 21 {
					suggest = suggestLevel3
				} else if fatRate > 21 && fatRate <= 26 {
					suggest = suggestLevel4
				} else if fatRate > 26 {
					suggest = suggestLevel5
				}

			} else if age >= 40 && age <= 59 {
				if fatRate >= 0 && fatRate <= 11 {
					suggest = suggestLevel1
				} else if fatRate > 11 && fatRate <= 17 {
					suggest = suggestLevel2
				} else if fatRate > 17 && fatRate <= 22 {
					suggest = suggestLevel3
				} else if fatRate > 22 && fatRate <= 27 {
					suggest = suggestLevel4
				} else if fatRate > 27 {
					suggest = suggestLevel5
				}

			} else if age >= 60 {
				if fatRate >= 0 && fatRate <= 13 {
					suggest = suggestLevel1
				} else if fatRate > 13 && fatRate <= 19 {
					suggest = suggestLevel2
				} else if fatRate > 19 && fatRate <= 24 {
					suggest = suggestLevel3
				} else if fatRate > 24 && fatRate <= 29 {
					suggest = suggestLevel4
				} else if fatRate > 29 {
					suggest = suggestLevel5
				}

			}

		} else if genderType == 0 {
			// 女 体脂率 建议
			if age >= 18 && age <= 39 {
				if fatRate >=0 && fatRate <= 20 {
					suggest = suggestLevel1
				} else if fatRate > 20 && fatRate <= 27 {
					suggest = suggestLevel2
				} else if fatRate > 27 && fatRate <= 34 {
					suggest = suggestLevel3
				} else if fatRate > 34 && fatRate <= 39 {
					suggest = suggestLevel4
				} else if fatRate > 39 {
					suggest = suggestLevel5
				}

			} else if age >= 40 && age <= 59 {
				if fatRate >= 0 && fatRate <= 21 {
					suggest = suggestLevel1
				} else if fatRate > 21 && fatRate <= 28 {
					suggest = suggestLevel2
				} else if fatRate > 28 && fatRate <= 35 {
					suggest = suggestLevel3
				} else if fatRate > 35 && fatRate <= 40 {
					suggest = suggestLevel4
				} else if fatRate > 40 {
					suggest = suggestLevel5
				}

			} else if age >= 60 {
				if fatRate >= 0 && fatRate <= 22 {
					suggest = suggestLevel1
				} else if fatRate > 22 && fatRate <= 29 {
					suggest = suggestLevel2
				} else if fatRate > 29 && fatRate <= 36 {
					suggest = suggestLevel3
				} else if fatRate > 36 && fatRate <= 41 {
					suggest = suggestLevel4
				} else if fatRate > 41 {
					suggest = suggestLevel5
				}

			}
		}

		var fatRatePercent = fatRate * 100
		fmt.Println("姓名:", name, " 性别:", gender, " 身高(m):", stature, " 体重(kg):", weight, " 年龄:", age,
			" genderType:", genderType, " bmi:", bmi, " 体脂率:", fatRatePercent, "% 建议:", suggest)
		calculateNum++

	}

	averageFatRate = fatRateSum / float64(calculateNum)

	fmt.Println("体脂率计算器 录入人数:", calculateNum, " 平均体脂率:", averageFatRate, " 总体脂率:", fatRateSum)

}
