# howto

<p align="center">
  <br>
  <img src=".github/images/demo.gif" width="600" alt="walk demo">
  <br>
</p>

**Howto** is a terminal helper which queries OpenAI API and inserts the result into the current terminal input.

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

Write a command in terminal and press `ctrl+g` to send current command to OpenAI API.
LLM response will be inserted into the current input. You can **modify** the returned command,
before executing it.

Recall previous command from history and to adjust the prompt.

## Related

- [walk](https://github.com/antonmedv/walk) – terminal file manager
- [fx](https://github.com/antonmedv/fx) – terminal JSON viewer

## License

[MIT](LICENSE)
