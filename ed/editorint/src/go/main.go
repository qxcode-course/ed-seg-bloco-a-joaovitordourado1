package main

import (
    "fmt"

    "github.com/gdamore/tcell/v2"
)

type Editor struct {
    lines  *List[*List[rune]]
    line   *Node[*List[rune]]
    cursor *Node[rune]
    screen tcell.Screen
    style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
    e.cursor = e.line.Value.Insert(e.cursor, r)
    e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
    if e.cursor != e.line.Value.Front() {
        e.cursor = e.cursor.Prev()
        return
    }
    if e.line != e.lines.Front() {
        e.line = e.line.Prev()
        e.cursor = e.line.Value.End()
    }
}

func (e *Editor) KeyEnter() {
    novaLinha := NewList[rune]()
    
    if e.cursor == e.line.Value.End() {
        e.lines.Insert(e.line.Next(), novaLinha)
        e.line = e.line.Next()
        e.cursor = e.line.Value.Front()
        return
    }
    for no := e.cursor; no != e.line.Value.End(); {
        novaLinha.PushBack(no.Value)
        no = e.line.Value.Erase(no)
    }
    e.lines.Insert(e.line.Next(), novaLinha)
    e.line = e.line.Next()
    e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
    if e.cursor != e.line.Value.End() {
        e.cursor = e.cursor.Next()
        return
    }
    if e.line.Next() != e.lines.End() {
        e.line = e.line.Next()
        e.cursor = e.line.Value.Front()
    }
}

func (e *Editor) KeyUp() {
    if e.line != e.lines.Front() {
        k := e.line.Value.IndexOf(e.cursor)
        e.line = e.line.Prev()
        m := e.line.Value.IndexOf(e.line.Value.End())
        if m < k {
            e.cursor = e.line.Value.End()
        } else {
            e.cursor = e.line.Value.Front()
            for j := 0; j < k; j++ {
                e.cursor = e.cursor.Next()
            }
        }
    }
}

func (e *Editor) KeyDown() {
    if e.line.Next() != e.lines.End() {
        k := e.line.Value.IndexOf(e.cursor)
        e.line = e.line.Next()
        m := e.line.Value.IndexOf(e.line.Value.End())
        if m < k {
            e.cursor = e.line.Value.End()
        } else {
            e.cursor = e.line.Value.Front()
            for j := 0; j < k; j++ {
                e.cursor = e.cursor.Next()
            }
        }
    }
}

func (e *Editor) KeyBackspace() {
    if e.cursor != e.line.Value.Front() {
        e.cursor = e.cursor.Prev()
        e.cursor = e.line.Value.Erase(e.cursor)
        return
    }
    
    if e.line != e.lines.Front() {
        linhaAnterior := e.line.Prev()
        alvoCursor := linhaAnterior.Value.End()
        primeiroMovido := false

        for no := e.line.Value.Front(); no != e.line.Value.End(); {
            linhaAnterior.Value.PushBack(no.Value)
            if !primeiroMovido {
                alvoCursor = linhaAnterior.Value.Back()
                primeiroMovido = true
            }
            no = e.line.Value.Erase(no)
        }
        
        e.lines.Erase(e.line)
        e.line = linhaAnterior
        e.cursor = alvoCursor
    }
}

func (e *Editor) KeyDelete() {
    if e.cursor != e.line.Value.End() {
        e.cursor = e.line.Value.Erase(e.cursor)
        return
    }
    if e.line.Next() != e.lines.End() {
        proxima := e.line.Next()
        ultimo := e.line.Value.End()

        for no := proxima.Value.Front(); no != proxima.Value.End(); {
            e.line.Value.PushBack(no.Value)
            no = proxima.Value.Erase(no)
        }
        e.lines.Erase(proxima)
        e.cursor = ultimo
    }
}

func main() {
    editor := NewEditor()
    editor.Draw()
    editor.MainLoop()
    defer editor.screen.Fini()
}

func (e *Editor) MainLoop() {
    for {
        ev := e.screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            switch ev.Key() {
            case tcell.KeyEsc, tcell.KeyCtrlC:
                return
            case tcell.KeyEnter:
                e.KeyEnter()
            case tcell.KeyLeft:
                e.KeyLeft()
            case tcell.KeyRight:
                e.KeyRight()
            case tcell.KeyUp:
                e.KeyUp()
            case tcell.KeyDown:
                e.KeyDown()
            case tcell.KeyBackspace, tcell.KeyBackspace2:
                e.KeyBackspace()
            case tcell.KeyDelete:
                e.KeyDelete()
            default:
                if ev.Rune() != 0 {
                    e.InsertChar(ev.Rune())
                }
            }
            e.Draw()
        case *tcell.EventResize:
            e.screen.Sync()
            e.Draw()
        }
    }
}

func NewEditor() *Editor {
    e := &Editor{}
    screen, err := tcell.NewScreen()
    if err != nil {
        fmt.Printf("erro ao criar a tela: %v", err)
    }
    if err := screen.Init(); err != nil {
        fmt.Printf("erro ao iniciar a tela: %v", err)
    }
    e.screen = screen
    e.lines = NewList[*List[rune]]()
    e.lines.PushBack(NewList[rune]())
    e.line = e.lines.Front()
    e.cursor = e.line.Value.Back()
    e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

    e.screen.SetStyle(e.style)
    e.screen.Clear()
    return e
}

func (e *Editor) Draw() {
    e.screen.Clear()
    x := 0
    y := 0
    for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
        for char := line.Value.Front(); ; char = char.Next() {
            data := char.Value
            if char == line.Value.End() {
                data = '⤶'
            }
            if data == ' ' {
                data = '·'
            }
            if char == e.cursor {
                e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
            } else {
                e.screen.SetContent(x, y, data, nil, e.style)
            }
            x++
            if char == line.Value.End() {
                break
            }
        }
        y++
        x = 0
    }
    e.screen.Show()
}