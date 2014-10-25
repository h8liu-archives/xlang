editorInit = ->
    editor = ace.edit("editor")

    editor.setTheme("ace/theme/tomorrow")
    editor.getSession().setMode("ace/mode/golang")
    editor.renderer.setShowGutter(false)
    editor.setHighlightActiveLine(false)
    editor.setShowFoldWidgets(false)
    editor.setDisplayIndentGuides(false)
    editor.setReadOnly(false)
    editor.setOptions({
        maxLines: Infinity,
        minLines: 20,
        fontFamily: "Consolas, Inconsolata, Monaco, \"Courier New\", Courier, monospace",
        fontSize: "13px",
    })
    editor.commands.removeCommands(["gotoline", "find"])

    prog = [
        'var x = 3',
        'var y = 4',
        '',
        'func main() {',
        '\tprintln(x + y)',
        '}',
    ].join('\n')

    editor.setValue(prog)
    editor.clearSelection()
    return

main = ->
    editorInit()

    $("#tokens").hide()
    $("#console").hide()

    $("#but-edit").click( (e) ->
        e.preventDefault()
        return
    )

    $("#but-tokens").click( (e) ->
        e.preventDefault()
        return
    )

    $("#but-console").click( (e) ->
        e.preventDefault()
        return
    )
    return

$(document).ready(main)
