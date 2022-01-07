package config

type MailConfig struct {
	Sender    string
	Receivers []string
	Password  string
	Smtphost  string
	Smtpport  int
}

type FlinkConfig struct {
	Endpoint string
}

type Config struct {
	Poll  int // seconds
	For   int // seconds
	Mail  MailConfig
	Flink FlinkConfig
}
