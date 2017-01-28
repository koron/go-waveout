package waveout

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zwaveout_windows.go waveout_windows.go

//sys   Open(handle *syscall.Handle, deviceID uint32, waveFormat *WaveFormatEx, callback uint32, inst uint32, flag uint32) (result MMRESULT) = winmm.waveOutOpen
//sys   Close(handle syscall.Handle) (result MMRESULT) = winmm.waveOutClose
//sys   GetVolume(handle syscall.Handle, volume *uint32) (result MMRESULT) = winmm.waveOutGetVolume
//sys   SetVolume(handle syscall.Handle, volume uint32) (result MMRESULT) = winmm.waveOutSetVolume
