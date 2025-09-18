package models

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	QuestionNumber int    `json:"question_number"`
	QuestionText   string `json:"question_text" gorm:"type:text"`
	Options        string `json:"options" gorm:"type:text"` // JSON string of options
}

type CustomerResult struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Gender       string    `json:"gender"`
	Answers      string    `json:"answers"`      // "ABCDE" format
	FinalResult  string    `json:"final_result"` // The Maestro, etc.
	TraitScores  string    `json:"trait_scores"` // JSON string
	CreatedAt    time.Time `json:"created_at"`
}

type QuizSubmission struct {
	Name    string   `json:"name" binding:"required"`
	Phone   string   `json:"phone"`
	Gender  string   `json:"gender" binding:"required"`
	Answers []string `json:"answers" binding:"required"`
}

type QuizResult struct {
	Result      string            `json:"result"`
	Description string            `json:"description"`
	Scores      map[string]int    `json:"scores"`
}

// Seed questions into database
func SeedQuestions(db *gorm.DB) {
	var count int64
	db.Model(&Question{}).Count(&count)
	if count > 0 {
		return // Questions already exist
	}

	questions := []Question{
		{
			QuestionNumber: 1,
			QuestionText:   "Khi thấy mình chìm vào thế giới riêng tư, bạn thường...",
			Options: `{
				"A": "Cảm nhận rõ sự kiên định, tự chủ và giá trị cá nhân.",
				"B": "Thường viết hoặc ghi chú lại các xúc cảm, hình ảnh trong đầu.",
				"C": "Thích một khoảng yên lặng, lắng nghe những chi tiết nhỏ trong không gian.",
				"D": "Bỗng thấy năng lượng của mình lan tỏa đến những người xung quanh.",
				"E": "Phóng tâm trí vào những cảm xúc và màu sắc mới mẻ."
			}`,
		},
		{
			QuestionNumber: 2,
			QuestionText:   "Trong đám đông, bạn hay đóng vai trò...",
			Options: `{
				"A": "Đem đến những ý tưởng mới, giúp mọi người nhìn nhận khác đi.",
				"B": "Để lại dấu ấn rõ ràng bằng phong thái vững chãi, không cần thể hiện quá lời.",
				"C": "Là người chia sẻ sâu sắc và lắng nghe những câu chuyện tâm trạng.",
				"D": "Người nhẹ nhàng, không muốn là trung tâm, nhưng để người khác cảm thấy dễ chịu.",
				"E": "Đứng ở vị trí dẫn dắt, thu hút sự chú ý và truyền cảm hứng."
			}`,
		},
		{
			QuestionNumber: 3,
			QuestionText:   "Khi phải đối mặt với thay đổi, bạn sẽ...",
			Options: `{
				"A": "Chủ động đón nhận, sẵn sàng dẫn dắt để mọi thứ trở thành cơ hội.",
				"B": "Xem thay đổi như chất liệu để sáng tạo, đổi mới cảm xúc.",
				"C": "Tìm cách thích nghi với mọi thứ một cách tinh tế, không ồn ào.",
				"D": "Cảm nhận sâu sắc, ghi lại những suy tư và góc nhìn cá nhân.",
				"E": "Kiên quyết khẳng định lập trường, chứng tỏ giá trị bản thân không thay đổi."
			}`,
		},
		{
			QuestionNumber: 4,
			QuestionText:   "Khi kết nối với mọi người, cách bạn tạo ấn tượng nhất là...",
			Options: `{
				"A": "Để lại dấu ấn rõ rệt qua từng cử chỉ, lời nói, không cần phô trương.",
				"B": "Mang đến cái nhìn khác biệt, giúp mọi người mở rộng góc nhìn và cảm nhận.",
				"C": "Kết nối bằng cảm xúc, chia sẻ mọi suy tư thầm kín nhất.",
				"D": "Tỏa ra nguồn lực, giúp người khác tin tưởng và cảm thấy an tâm khi ở bên.",
				"E": "Sự thanh lịch, nhẹ nhàng, thấu hiểu từng người một cách tinh tế."
			}`,
		},
		{
			QuestionNumber: 5,
			QuestionText:   "Nếu cuộc đời là một tác phẩm nghệ thuật, bạn muốn nó được ghi nhớ như...",
			Options: `{
				"A": "Một bức tranh đa sắc, nơi mỗi vệt màu là một cảm xúc bất ngờ.",
				"B": "Một áng thơ sâu lắng, kể lại những rung động và suy tư thầm kín.",
				"C": "Một bản giao hưởng hùng tráng, có khả năng truyền cảm hứng và dẫn dắt.",
				"D": "Một vũ điệu tinh tế, không cần lời nhưng vẫn chạm đến mọi cảm xúc.",
				"E": "Một công trình kiến trúc vững chãi, ghi lại dấu ấn vượt thời gian."
			}`,
		},
	}

	for _, question := range questions {
		db.Create(&question)
	}
}