package main

import (
	"context"
	"fmt"
	"os"

	"ice-breaker/pkg/helpers"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

func main() {
	_ = helpers.LoadDotEnv(".env")
	openai_api_key, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		panic("OPENAI_API_KEY environment variable not set")
	}

	var opts []openai.Option
	opts = append(opts, openai.WithToken(openai_api_key))
	opts = append(opts, openai.WithModel("gpt-4-turbo"))

	llm, err := openai.New(opts...)
	if err != nil {
		panic(err)
	}

	information := `
	Elon Reeve Musk (/ˈiːlɒn/ EE-lon; born June 28, 1971) is a businessman and entrepreneur known for his leadership of Tesla, SpaceX, Twitter, and xAI. Musk has been the wealthiest person in the world since 2021; as of October 2025, Forbes estimates his net worth to be US$500 billion.
	Born into a wealthy family in Pretoria, South Africa, Musk emigrated in 1989 to Canada; he had obtained Canadian citizenship at birth through his Canadian-born mother. He received bachelor's degrees in 1997 from the University of Pennsylvania in Philadelphia, United States, before moving to California to pursue business ventures. In 1995, Musk co-founded the software company Zip2. Following its sale in 1999, he co-founded X.com, an online payment company that later merged to form PayPal, which was acquired by eBay in 2002. That year, Musk also became an American citizen.
	In 2002, Musk founded the space technology company SpaceX, becoming its CEO and chief engineer; the company has since led innovations in reusable rockets and commercial spaceflight. Musk joined the automaker Tesla as an early investor in 2004 and became its CEO and product architect in 2008; it has since become a leader in electric vehicles. In 2015, he co-founded OpenAI to advance artificial intelligence (AI) research, but later left, growing discontent with the organization's direction and their leadership in the AI boom in the 2020s led him to establish xAI. In 2022, he acquired the social network Twitter, implementing significant changes, and rebranding it as X in 2023. His other businesses include the neurotechnology company Neuralink, which he co-founded in 2016, and the tunneling company the Boring Company, which he founded in 2017.
	Musk was the largest donor in the 2024 U.S. presidential election, and is a supporter of global far-right figures, causes, and political parties. In early 2025, he served as senior advisor to United States president Donald Trump and as the de facto head of DOGE. After a public feud with Trump, Musk left the Trump administration and returned to his technology companies.
	Musk's political activities, views, and statements have made him a polarizing figure, especially following the COVID-19 pandemic. He has been criticized for making unscientific and misleading statements, including COVID-19 misinformation and promoting conspiracy theories, and affirming antisemitic, racist, and transphobic comments. His acquisition of Twitter was controversial due to a subsequent increase in hate speech and the spread of misinformation on the service. His role in the second Trump administration attracted public backlash, particularly in response to DOGE.
	`
	summary_template := `
	given the information {{.information}} regarding a person, I want you to create:
	1. A short summary about the person in 2-3 sentences.
	2. Two fun facts about the person.
	3. A question to ask the person to break the ice.
	`

	summary_prompt_template := prompts.NewPromptTemplate(
		summary_template,
		[]string{"information"},
	)

	llmchain := chains.NewLLMChain(llm, summary_prompt_template)
	summaryOutputValues, err := chains.Call(context.Background(), llmchain, map[string]any{
		"information": information,
	})
	if err != nil {
		panic(err)
	}

	summary, ok := summaryOutputValues[llmchain.OutputKey].(string)
	if !ok {
		panic(fmt.Errorf("invalid chain return"))
	}
	fmt.Println(summary)
}
