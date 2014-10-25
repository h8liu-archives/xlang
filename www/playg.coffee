
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
        minLines: 20,
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
    
    session = editor.getSession()
    # session.addMarker(new Range(4, 1, 5, 3), "ace_selected-word", "text")

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

    # $("#tokens").hide()
    updateTokens()

    $("#console").hide()

    $("#but-edit").click( (e) ->
        e.preventDefault()
        return
    )

    $("#but-tokens").click( (e) ->
        e.preventDefault()
        updateTokens()
        return
    )

    xlang.editor.getSession().on("change", ->
        updateTokens()
        return
    )

    $("#but-console").click( (e) ->
        e.preventDefault()
        return
    )
    return

$(document).ready(main)
