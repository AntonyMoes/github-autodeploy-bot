# github-autodeploy-bot
Go bot based on Github webhookes

- Uses webhookes.json as config file
- If configured right (using bash scripts) can even deploy new versions of itself!

### webhookes.json example


```json
{
  "port": ":8080",
  "url": "/webhook",
  "repos":
  {
    "AntonyMoes/github-autodeploy-bot":
    {
      "master": "./examples/redeploy_bot.sh"
    }
  }
}
```
