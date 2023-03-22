REST API mode
==

Application in REST API mode provides single endpoint where API user can pass some text or question and receive ChatGPT 
answer as a response. In order to preserver communication context each mesage except first one maintain context handle 
identificator. 

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
