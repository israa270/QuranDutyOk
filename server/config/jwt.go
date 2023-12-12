package config

// JWT struct
type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt sign name
	ExpiresTime string  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // expiration
	BufferTime  string  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // buffer time
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                  // sign origin
}
