package main

//todo This block about open two and more windows in Goroutines
//probably this will need for external apps
//in the future, the logic will be redone as it is just a stub

/*
var (
	Sync sync.WaitGroup
)

func GoroutineAddWindow() {
	Sync.Add(1)
	go AddWindow(&Sync)
	Sync.Wait()
}

func AddWindow(w *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	win := winc.NewForm(nil)
	win.SetSize(900, 400)
	win.Center()
	win.Show()
	win.OnClose().Bind(wndOnClose)

	winc.RunMainLoop()
	defer w.Done()
}

func GoroutineStartWindow() {
	Sync.Add(1)
	go StartWindow(&Sync)
	Sync.Wait()
}

func StartWindow(w *sync.WaitGroup) {
	window := winc.NewForm(nil)
	window.SetMaxSize(1440, 900)
	window.SetSize(800, 400)

	//parent = controller (parameters) "win.NewEdit"
	//text form
	edt := winc.NewEdit(window)
	edt.SetMaxSize(1440, 900)
	edt.SetSize(800, 400)
	edt.SetPos(0, 30)

	//keyEnter := winc.KeyAccept
	//button
	btn := winc.NewPushButton(window)
	btn.SetText("show/hide")
	btn.SetPos(0, 0)
	btn.SetSize(60, 30)
	btn.OnClick().Bind(func(e *winc.Event) {
		if edt.Visible() {
			edt.Hide()
		} else {
			edt.Show()
		}
	})

	window.Center()
	window.Show()
	window.OnClose().Bind(wndOnClose)

	winc.RunMainLoop() // must call to start event loop
	defer w.Done()
}*/
