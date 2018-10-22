package main

type mockEndCmd struct {
	shutdownCMD []string
	syncCMD     []string
}

func (m *mockEndCmd) Run() {
}
