editorInit = ->
    console.log("here")
    Range = ace.require("ace/range").Range
    editor = ace.edit("editor")
    xlang.editor = editor
    console.log(editor.getValue())
    session = editor.getSession()

    editor.setTheme("ace/theme/tomorrow")
    session.setMode("ace/mode/golang")
    editor.renderer.setShowGutter(false)
    editor.setHighlightActiveLine(false)
    editor.setShowFoldWidgets(false)
    editor.setDisplayIndentGuides(false)
    # editor.setReadOnly(false)
    ff = "Consolas, Inconsolata, Monaco, \"Courier New\", Courier, monospace"
    editor.setOptions({
        maxLines: Infinity,
        minLines: 10,
        fontFamily: ff,
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

exampleInit = ->
    examples = ['3p4']

    ul = $('<ul id="examples"/>')
    for f in examples
        li = $('<li><a href="#">' + f + '</a></li>')
        li.find('a').click( (e) ->
            e.preventDefault()
            # TODO: load the file
            console.log("load "+f)
            return
        )
        ul.append(li)

    $("div#filelist").append(ul)
    return

updateTokens = ->
    code = xlang.editor.getValue()
    parsed = xlang.parseTokens("", code)
    $("#tokens").html(parsed)
    return

main = ->
    editorInit()
    exampleInit()
    updateTokens()
    xlang.editor.getSession().on("change", ->
        updateTokens()
        return
    )
    return

$(document).ready(main)
