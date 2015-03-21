# Slacker News: A Slack Integration News Service
Read current stories from Hacker News, the BBC, and more right from within Slack.
![Demo gif of usage](http://i.imgur.com/Tt8SDvu.gif)

The following sites are supported:
- Product Hunt
- Hacker News
- BBC
- Vice News
- FiveThirtyEight

## Installation

Set up [a new slash command](my.slack.com/services/new/slash-commands).

1. The first thing to assign is a name for the command. 
   
   Select something memorable like `/news` and then click `"Add Slash Command Integration"`

2. Change the `Url` to be `www.slckr-nws.heroku.com`

3. Change the `Method` to be GET from the default POST.

Voila! You are set up and ready to receive new and trending news in Slack. The remaining fields you can set to your preference or leave as they are. 
