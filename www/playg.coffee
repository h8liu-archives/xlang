editorInit = ->
    Range = ace.require("ace/range").Range
    editor = ace.edit("editor")
    xlang.editor = editor
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
            # console.log("load file: "+f)
            $.ajax("tests/" + f + ".x", {
                success: (dat) ->
                    xlang.editor.setValue(dat)
                    xlang.editor.clearSelection()
                    return
            })
            return
        )
        ul.append(li)

    $("div#filelist").append(ul)
    return

updateTokens = ->
    code = xlang.editor.getValue()
    tokens = xlang.parseTokens("", code)
    $("#tokens").html(tokens)
    parsed = xlang.parse("test.x", code)
    $("#console").html(parsed.errs)

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
