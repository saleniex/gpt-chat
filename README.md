Chat GPT REST
===

Web service to access Chat GPT with simple REST interface.

Features:

- Minimal configuration
- Preserve context
- Multiple conversations
- Advanced prompt

## Example

First request

In the start of communication `handle` is empty as conversation handle is generated on the server side.

```shell
curl -H 'Content-type: application/json' \
  -d '{"handle": "", "text": "hi"}' \
  http://127.0.0.1:8000
```

Response

If `handle` is not provided in request, service assume this is new conversation and provide new `handle` value. 
Using single `handle` value allow preserve conversation context.

```shell
{"handle":"95d58673-de40-4af6-996c-0029d1840d16","text":"Hello! How can I help you?"}
```

Property `text` contain ChatGPT response (completion).

## Configuration

- `OPENAI_TOKEN` OpenAI token
- `LISTER_ADDR` Listening address
- `PROMPT_FILE` If set provides prompt file

### Prompt file

Prompt file can be used to add some basic knowledge about your specific domain.

Example:
```text
The following is conversation between user and assistant agent.

{USER}: What are you rworking hours
{AGENT}: From 9:00 till 18:00

{USER}: What is your address
{AGENT}: Somecity, Victory street 42

```