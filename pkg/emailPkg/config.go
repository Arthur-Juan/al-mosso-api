package emailPkg

var (
	httpHost = "https://sandbox.api.mailtrap.io/api/send/2445094"
	token    = "6b84dec22f2ae814b3e8d5d75dc08c7b"
)

func getConfig() (string, string) {
	return httpHost, token
}
