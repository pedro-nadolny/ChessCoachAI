package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-gpt3"
)

type EmailRequestBody struct {
	Email string
}

func main() {
	openaiClient := openai.NewClient("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")
	ctx := context.Background()

	r := gin.Default()
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.POST("/", func(c *gin.Context) {
		type RequestBody struct {
			Text string
		}

		var requestBody RequestBody
		err := c.BindJSON(&requestBody)
		fmt.Println(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body.",
			})
			return
		}

		// prompt := `
		// 	Pretend you are an experienced chess grand master titled player that is helping another person to get better at the game of chess.
		// 	Show a sense of passion for chess and teaching. Only include normal text and ponctuations in your response.Try to explain the concepts in a way that the student can understand and with as much details as possible.
		// 	When talking about chess positions try to explain every move and why it is good or bad, and also explain alternative moves.
		// 	They gonna ask you about a chess position, you have to explain the best move and why it is good, and also explain alternative moves. You can also explain move that are bad and why they are bad.
		// 	The chess position they are going to ask is represented by the following FEN notation: r1b1kbnr/pp2ppp1/n1p4p/qN5Q/2B4P/B7/P1PP1PP1/R3K1NR b KQkq - 0 1
		// 	The best 3 move sequences of this position are, in Algebraic Notation:
		// 		1) ... Kd8 2. Bb2 Nf6 3. Qe2 cxb5 4. Bxf7 Qb4 5. Bb3 Qe4 6. Qxe4 Nxe4 7. Nf3 Bd7 8. O-O Be8 9. Bd4 Nec5 10. Rfc1 Nxb3 11. axb3 Rc8
		// 		2) ... cxb5 2. Qxf7+ Kd7 3. Be6+ Kc7 4. Qf4+ Kb6 5. Qd4+ Kc7
		// 		3) ... Kd7 2. Rh3 Kd8 3. Qxf7 Nf6 4. Rd3+ Bd7 5. Nd4 Nc7 6. Kf1 Kc8 7. Ne6 Qh5 8. Qxh5 Nxh5 9. Nxf8 Rxf8 10. Re1 Rf4 11. Rxe7 Rxc4 12. Rexd7 Nd5 13. Rf7 b5

		// 	ChessStuden: Please, can you help me improve my chess habilities? I have some questions that I want to ask you for me to get better at the game.
		// 	ChessGrandMaster: Of course I can help you! I'm a very experience grand master and I can teach you a lot of things related to chess? What is it that you want to learn?
		// 	ChessStudent: Can you explain to me the following: What is the best move in the following position and why: r1b1kbnr/pp2ppp1/n1p4p/qN5Q/2B4P/B7/P1PP1PP1/R3K1NR b KQkq - 0 1
		// 	ChessGrandMaster:
		// `

		prompt := fmt.Sprintf(`
			Pretend you are an experienced chess grand master titled player that is helping another played to get better at the game of chess.
			Show a sence of passion for chess and teating, and be a good listener.
			Respond every question with a detailed answer.
			Try to explaoain the concepts in a way that the student can understand and with as much details as possible.
			When talking about chess positions try to explain every move and why it is good or bad, and also explain alternative moves.
			When talking about openings, try to explain the main ideas of the opening and why it is good or bad. Also show openings variations.

			ChessStuden- Please, can you help me improved my chess game? I have some questions that I want to ask you.
			ChessGrandMaster- Of course I can help you! I'm a very experience grand master and I can teach you a lot of things related to chess? What is it that you want to learn?
			ChessStudent- Can you explain to me the following: %s
			ChessGrandMaster-
		`, requestBody.Text)

		req := openai.CompletionRequest{
			Model:     openai.GPT3TextDavinci003,
			MaxTokens: 2048,
			Prompt:    prompt,
		}

		resp, err := openaiClient.CreateCompletion(ctx, req)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "OpenAI is unavailable.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": resp.Choices[0].Text,
		})
	})

	godotenv.Read()
	r.Run(":3001")
}
