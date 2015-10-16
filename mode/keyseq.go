package mode

type Operator interface {
	DoKeyEsc()
	DoKeyArrowLeft()
	DoKeyCtrlB()
	DoKeyArrowRight()
	DoKeyCtrlF()
	DoKeyBackspace()
	DoKeyDelete()
	DoKeyTab()
	DoKeyCtrlX()
	DoKeyCtrlW()
	DoKeyCtrlL()
	DoKeyCtrlV()
	DoKeyCtrlD()
	DoKeyCtrlA()
	DoKeyCtrlR()
	DoKeySpace()
	DoEnter()
	DoChar(rune)
}

type Nothing struct{}

func (n Nothing) DoKeyEsc()        {}
func (n Nothing) DoKeyArrowLeft()  {}
func (n Nothing) DoKeyCtrlB()      {}
func (n Nothing) DoKeyArrowRight() {}
func (n Nothing) DoKeyCtrlF()      {}
func (n Nothing) DoKeyBackspace()  {}
func (n Nothing) DoKeyDelete()     {}
func (n Nothing) DoKeyTab()        {}
func (n Nothing) DoKeyCtrlX()      {}
func (n Nothing) DoKeyCtrlW()      {}
func (n Nothing) DoKeyCtrlL()      {}
func (n Nothing) DoKeyCtrlV()      {}
func (n Nothing) DoKeyCtrlD()      {}
func (n Nothing) DoKeyCtrlA()      {}
func (n Nothing) DoKeyCtrlR()      {}
func (n Nothing) DoKeySpace()      {}
func (n Nothing) DoEnter()         {}
func (n Nothing) DoChar(r rune)    {}
