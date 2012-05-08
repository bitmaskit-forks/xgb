package xgb

/*
	This file was generated by dpms.xml on May 7 2012 11:34:25pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// DpmsInit must be called before using the DPMS extension.
func (c *Conn) DpmsInit() error {
	reply, err := c.QueryExtension(4, "DPMS").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return errorf("No extension named DPMS could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["DPMS"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["DPMS"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range newExtErrorFuncs["DPMS"] {
		newErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["DPMS"] = make(map[int]newEventFun)
	newExtErrorFuncs["DPMS"] = make(map[int]newErrorFun)
}

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Id'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

const (
	DpmsDPMSModeOn      = 0
	DpmsDPMSModeStandby = 1
	DpmsDPMSModeSuspend = 2
	DpmsDPMSModeOff     = 3
)

// Request DpmsGetVersion
// size: 8
type DpmsGetVersionCookie struct {
	*cookie
}

func (c *Conn) DpmsGetVersion(ClientMajorVersion uint16, ClientMinorVersion uint16) DpmsGetVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.dpmsGetVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return DpmsGetVersionCookie{cookie}
}

func (c *Conn) DpmsGetVersionUnchecked(ClientMajorVersion uint16, ClientMinorVersion uint16) DpmsGetVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.dpmsGetVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return DpmsGetVersionCookie{cookie}
}

// Request reply for DpmsGetVersion
// size: 12
type DpmsGetVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

// Waits and reads reply data from request DpmsGetVersion
func (cook DpmsGetVersionCookie) Reply() (*DpmsGetVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return dpmsGetVersionReply(buf), nil
}

// Read reply into structure from buffer for DpmsGetVersion
func dpmsGetVersionReply(buf []byte) *DpmsGetVersionReply {
	v := new(DpmsGetVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = Get16(buf[b:])
	b += 2

	v.ServerMinorVersion = Get16(buf[b:])
	b += 2

	return v
}

func (cook DpmsGetVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsGetVersion
func (c *Conn) dpmsGetVersionRequest(ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put16(buf[b:], ClientMajorVersion)
	b += 2

	Put16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}

// Request DpmsCapable
// size: 4
type DpmsCapableCookie struct {
	*cookie
}

func (c *Conn) DpmsCapable() DpmsCapableCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.dpmsCapableRequest(), cookie)
	return DpmsCapableCookie{cookie}
}

func (c *Conn) DpmsCapableUnchecked() DpmsCapableCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.dpmsCapableRequest(), cookie)
	return DpmsCapableCookie{cookie}
}

// Request reply for DpmsCapable
// size: 32
type DpmsCapableReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Capable bool
	// padding: 23 bytes
}

// Waits and reads reply data from request DpmsCapable
func (cook DpmsCapableCookie) Reply() (*DpmsCapableReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return dpmsCapableReply(buf), nil
}

// Read reply into structure from buffer for DpmsCapable
func dpmsCapableReply(buf []byte) *DpmsCapableReply {
	v := new(DpmsCapableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	if buf[b] == 1 {
		v.Capable = true
	} else {
		v.Capable = false
	}
	b += 1

	b += 23 // padding

	return v
}

func (cook DpmsCapableCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsCapable
func (c *Conn) dpmsCapableRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request DpmsGetTimeouts
// size: 4
type DpmsGetTimeoutsCookie struct {
	*cookie
}

func (c *Conn) DpmsGetTimeouts() DpmsGetTimeoutsCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.dpmsGetTimeoutsRequest(), cookie)
	return DpmsGetTimeoutsCookie{cookie}
}

func (c *Conn) DpmsGetTimeoutsUnchecked() DpmsGetTimeoutsCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.dpmsGetTimeoutsRequest(), cookie)
	return DpmsGetTimeoutsCookie{cookie}
}

// Request reply for DpmsGetTimeouts
// size: 32
type DpmsGetTimeoutsReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	StandbyTimeout uint16
	SuspendTimeout uint16
	OffTimeout     uint16
	// padding: 18 bytes
}

// Waits and reads reply data from request DpmsGetTimeouts
func (cook DpmsGetTimeoutsCookie) Reply() (*DpmsGetTimeoutsReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return dpmsGetTimeoutsReply(buf), nil
}

// Read reply into structure from buffer for DpmsGetTimeouts
func dpmsGetTimeoutsReply(buf []byte) *DpmsGetTimeoutsReply {
	v := new(DpmsGetTimeoutsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.StandbyTimeout = Get16(buf[b:])
	b += 2

	v.SuspendTimeout = Get16(buf[b:])
	b += 2

	v.OffTimeout = Get16(buf[b:])
	b += 2

	b += 18 // padding

	return v
}

func (cook DpmsGetTimeoutsCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsGetTimeouts
func (c *Conn) dpmsGetTimeoutsRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request DpmsSetTimeouts
// size: 12
type DpmsSetTimeoutsCookie struct {
	*cookie
}

// Write request to wire for DpmsSetTimeouts
func (c *Conn) DpmsSetTimeouts(StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) DpmsSetTimeoutsCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.dpmsSetTimeoutsRequest(StandbyTimeout, SuspendTimeout, OffTimeout), cookie)
	return DpmsSetTimeoutsCookie{cookie}
}

func (c *Conn) DpmsSetTimeoutsChecked(StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) DpmsSetTimeoutsCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.dpmsSetTimeoutsRequest(StandbyTimeout, SuspendTimeout, OffTimeout), cookie)
	return DpmsSetTimeoutsCookie{cookie}
}

func (cook DpmsSetTimeoutsCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsSetTimeouts
func (c *Conn) dpmsSetTimeoutsRequest(StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put16(buf[b:], StandbyTimeout)
	b += 2

	Put16(buf[b:], SuspendTimeout)
	b += 2

	Put16(buf[b:], OffTimeout)
	b += 2

	return buf
}

// Request DpmsEnable
// size: 4
type DpmsEnableCookie struct {
	*cookie
}

// Write request to wire for DpmsEnable
func (c *Conn) DpmsEnable() DpmsEnableCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.dpmsEnableRequest(), cookie)
	return DpmsEnableCookie{cookie}
}

func (c *Conn) DpmsEnableChecked() DpmsEnableCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.dpmsEnableRequest(), cookie)
	return DpmsEnableCookie{cookie}
}

func (cook DpmsEnableCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsEnable
func (c *Conn) dpmsEnableRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request DpmsDisable
// size: 4
type DpmsDisableCookie struct {
	*cookie
}

// Write request to wire for DpmsDisable
func (c *Conn) DpmsDisable() DpmsDisableCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.dpmsDisableRequest(), cookie)
	return DpmsDisableCookie{cookie}
}

func (c *Conn) DpmsDisableChecked() DpmsDisableCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.dpmsDisableRequest(), cookie)
	return DpmsDisableCookie{cookie}
}

func (cook DpmsDisableCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsDisable
func (c *Conn) dpmsDisableRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request DpmsForceLevel
// size: 8
type DpmsForceLevelCookie struct {
	*cookie
}

// Write request to wire for DpmsForceLevel
func (c *Conn) DpmsForceLevel(PowerLevel uint16) DpmsForceLevelCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.dpmsForceLevelRequest(PowerLevel), cookie)
	return DpmsForceLevelCookie{cookie}
}

func (c *Conn) DpmsForceLevelChecked(PowerLevel uint16) DpmsForceLevelCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.dpmsForceLevelRequest(PowerLevel), cookie)
	return DpmsForceLevelCookie{cookie}
}

func (cook DpmsForceLevelCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsForceLevel
func (c *Conn) dpmsForceLevelRequest(PowerLevel uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put16(buf[b:], PowerLevel)
	b += 2

	return buf
}

// Request DpmsInfo
// size: 4
type DpmsInfoCookie struct {
	*cookie
}

func (c *Conn) DpmsInfo() DpmsInfoCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.dpmsInfoRequest(), cookie)
	return DpmsInfoCookie{cookie}
}

func (c *Conn) DpmsInfoUnchecked() DpmsInfoCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.dpmsInfoRequest(), cookie)
	return DpmsInfoCookie{cookie}
}

// Request reply for DpmsInfo
// size: 32
type DpmsInfoReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	PowerLevel uint16
	State      bool
	// padding: 21 bytes
}

// Waits and reads reply data from request DpmsInfo
func (cook DpmsInfoCookie) Reply() (*DpmsInfoReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return dpmsInfoReply(buf), nil
}

// Read reply into structure from buffer for DpmsInfo
func dpmsInfoReply(buf []byte) *DpmsInfoReply {
	v := new(DpmsInfoReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.PowerLevel = Get16(buf[b:])
	b += 2

	if buf[b] == 1 {
		v.State = true
	} else {
		v.State = false
	}
	b += 1

	b += 21 // padding

	return v
}

func (cook DpmsInfoCookie) Check() error {
	return cook.check()
}

// Write request to wire for DpmsInfo
func (c *Conn) dpmsInfoRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["DPMS"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}
