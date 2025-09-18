package utils

import "encoding/json"

// Mapping table theo logic đã định
var answerMapping = map[int]map[string]string{
	1: {"A": "S", "B": "P", "C": "B", "D": "M", "E": "A"},
	2: {"A": "A", "B": "S", "C": "P", "D": "B", "E": "M"},
	3: {"A": "M", "B": "A", "C": "B", "D": "P", "E": "S"},
	4: {"A": "S", "B": "A", "C": "P", "D": "M", "E": "B"},
	5: {"A": "A", "B": "P", "C": "M", "D": "B", "E": "S"},
}

type ScoreResult struct {
	M int `json:"M"` // The Maestro
	S int `json:"S"` // The Sculptor
	P int `json:"P"` // The Poet
	B int `json:"B"` // The Ballerina
	A int `json:"A"` // The Artist
}

func CalculateResult(answers []string, gender string) (string, map[string]int, string) {
	scores := ScoreResult{}

	// Tính điểm cho từng trait
	for i, answer := range answers {
		if trait, exists := answerMapping[i+1][answer]; exists {
			switch trait {
			case "M":
				scores.M++
			case "S":
				scores.S++
			case "P":
				scores.P++
			case "B":
				scores.B++
			case "A":
				scores.A++
			}
		}
	}

	// Reset điểm theo giới tính
	if gender == "female" {
		scores.M = 0
		scores.S = 0
	} else if gender == "male" {
		scores.P = 0
		scores.B = 0
	}

	// Convert to map for JSON
	scoresMap := map[string]int{
		"M": scores.M,
		"S": scores.S,
		"P": scores.P,
		"B": scores.B,
		"A": scores.A,
	}

	// Tìm trait cao nhất
	result, description := findHighestTrait(scores, answers, gender)
	
	return result, scoresMap, description
}

func findHighestTrait(scores ScoreResult, answers []string, gender string) (string, string) {
	maxScore := 0
	traits := []struct {
		name  string
		score int
	}{
		{"M", scores.M},
		{"S", scores.S},
		{"P", scores.P},
		{"B", scores.B},
		{"A", scores.A},
	}

	// Tìm điểm cao nhất
	for _, trait := range traits {
		if trait.score > maxScore {
			maxScore = trait.score
		}
	}

	// Nếu tất cả điểm = 0, trả về A
	if maxScore == 0 {
		return getTraitName("A"), getTraitDescription("A")
	}

	// Lấy tất cả traits có điểm cao nhất
	highestTraits := []string{}
	for _, trait := range traits {
		if trait.score == maxScore {
			highestTraits = append(highestTraits, trait.name)
		}
	}

	// Nếu chỉ có 1 trait cao nhất
	if len(highestTraits) == 1 {
		return getTraitName(highestTraits[0]), getTraitDescription(highestTraits[0])
	}

	// Nếu có nhiều traits hòa, ưu tiên theo câu 5, câu 4
	for _, questionNum := range []int{5, 4} {
		if questionNum <= len(answers) {
			answerIndex := questionNum - 1
			trait := answerMapping[questionNum][answers[answerIndex]]
			
			// Kiểm tra trait này có trong danh sách cao nhất không
			for _, highTrait := range highestTraits {
				if trait == highTrait {
					return getTraitName(trait), getTraitDescription(trait)
				}
			}
		}
	}

	// Mặc định trả về A
	return getTraitName("A"), getTraitDescription("A")
}

func getTraitName(trait string) string {
	names := map[string]string{
		"M": "The Maestro",
		"S": "The Sculptor",
		"P": "The Poet",
		"B": "The Ballerina",
		"A": "The Artist",
	}
	return names[trait]
}

func getTraitDescription(trait string) string {
	descriptions := map[string]string{
		"M": "Bạn là người có phong thái mạnh mẽ, tự tin và có khả năng dẫn dắt. Mùi hương phù hợp với bạn thường có tính chất sang trọng, mạnh mẽ.",
		"S": "Bạn là người có tính cách vững chãi, kiên định và để lại ấn tượng sâu sắc. Mùi hương của bạn thường tinh tế nhưng đầy tính cách.",
		"P": "Bạn là người nhạy cảm, sâu sắc và có khả năng cảm nhận tinh tế. Mùi hương phù hợp với bạn thường mềm mại, thơ mộng.",
		"B": "Bạn là người thanh lịch, duyên dáng và có sức hút đặc biệt. Mùi hương của bạn thường nhẹ nhàng nhưng quyến rũ.",
		"A": "Bạn là người sáng tạo, độc đáo và có cái nhìn khác biệt. Mùi hương phù hợp với bạn thường độc đáo, không theo khuôn mẫu.",
	}
	return descriptions[trait]
}

func ScoresToJSON(scores map[string]int) string {
	jsonBytes, _ := json.Marshal(scores)
	return string(jsonBytes)
}