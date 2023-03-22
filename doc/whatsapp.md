Whatsapp bot mode
==

When application is used in Whatsapp bot mode the first thing you need to do is set up webhooks in order to receive 
notification about new messages sent by communication peer. When notification is received incoming text is passed to 
ChatGPT and response is sent back to user. So in order to be able to send messages to user you need authentication 
credentials.

A Whatsapp integration relays on communication from Facebook (notification on message arrival), application at first 
should be deployed with publicly accessible domain name. For thispurposu you can use either some serverless solutions 
(e.g. DigitalOcean Apps) or set up as docker container in any environment with docker support.

Detailed information on how to create [access tokens](https://developers.facebook.com/docs/whatsapp/business-management-api/get-started) 
and [set up webhooks](https://developers.facebook.com/docs/whatsapp/business-management-api/guides/set-up-webhooks) can 
be found in Meta for developers portal.

Here will follow steps to create required configuration:
1. In Meta For Developers open Business accounts
2. Click on "Create App"
3. When app is created add products Whatsapp and Webhooks
4. Create webhook with subscription of messages. Here you will need to set up application on publicly accessible host and use in webhook URL https://<yourdomain>/meta/webhooks
5. Use menu Whatsapp -> Getting started to test your integration and then Whatsapp -> Configuration to create permanent access token

Configuration of webhook and access token details you can find [here](configuration.md)
