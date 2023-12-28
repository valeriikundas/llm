import argparse

import replicate
import dotenv


def main():
    dotenv.load_dotenv()

    parser = argparse.ArgumentParser()
    parser.add_argument("prompt")
    args = parser.parse_args()

    deployment = replicate.deployments.get("valeriikundas/mistral-personal")
    prediction = deployment.predictions.create(
        input={
            "top_k": 50,
            "top_p": 0.9,
            "prompt": args.prompt,
            "temperature": 0.6,
            "max_new_tokens": 1024,
            "prompt_template": "<s>[INST] {prompt} [/INST] ",
            "presence_penalty": 0,
            "frequency_penalty": 0,
        }
    )
    # todo: show progress bar
    prediction.wait()
    print(prediction.output)


if __name__ == "__main__":
    main()
