# Telegram Music Bot in Go

An example bot using [gotgcalls].

## Setup

1. Install [the server].
2. Create a `.env` file with the following variables:

- `APP_ID` and `APP_HASH`, which are your Telegram app credentials obtained from
  [my.telegram.org/apps].
- `BOT_TOKEN`, which is your bot's token obtained from [@BotFather]

## Running

```bash
go run .
```

> Note that ou will need to login to your Telegram account once.

## Commands

| Command | Description                               |
| ------- | ----------------------------------------- |
| stream  | Streams or queues the replied audio file. |
| skip    | Skips the current stream.                 |
| stop    | Stops streaming and clears the queue.     |
| mute    | Mutes the stream audio.                   |
| unmute  | Unmutes the stream audio.                 |
| pause   | Pauses the stream.                        |
| resume  | Resumes the stream.                       |

> Note that the commands currently work in supergroups only.

[gotgcalls]: https://github.com/gotgcalls/tgcalls
[the server]: https://github.com/gotgcalls/server
[my.telegram.org/apps]: https://my.telegram.org/apps
[@BotFather]: https://t.me/BotFather
