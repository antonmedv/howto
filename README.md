# howto

<p align="center">
  <br>
  <img src=".github/images/demo.gif" width="600" alt="walk demo">
  <br>
</p>

**Howto** is a terminal helper which queries OpenAI API and inserts the result into the current terminal input.

Simply press <kbd>ctrl</kbd>+<kbd>g</kbd> to call **howto**. **Howto** replaces your command with a correct command from LLM.

## Install

```
go install github.com/antonmedv/howto@latest
```

Or download [prebuild binaries](https://github.com/antonmedv/howto/releases).

### Setup

Add a key binding to **.zshrc** or a similar config:

<table>
<tr>
  <th> Zsh </th>
</tr>
<tr>
<td>

```bash
bindkey -s "\C-g" "\C-ahowto \C-j"
```

</td>
</tr>
</table>

## Usage

Write a command in terminal and press <kbd>ctrl</kbd>+<kbd>g</kbd> to send current command to OpenAI API.
LLM response will be inserted into the current input. You can **modify** the returned command,
before executing it.

Recall previous command from history and to adjust the prompt.

## Examples

Use **howto** to list all container's hostnames:

<img src=".github/images/example-docker.gif" width="600" alt="howto example">

Use **howto** to convert a movie to mp4:

<img src=".github/images/example-ffmpeg.gif" width="600" alt="howto example">

## Related

- [walk](https://github.com/antonmedv/walk) – terminal file manager
- [fx](https://github.com/antonmedv/fx) – terminal JSON viewer
- [countdown](https://github.com/antonmedv/countdown) – terminal countdown timer


## License

[MIT](LICENSE)
