package captcha

import (
	"sync"
	"testing"
)

func TestNewCaptchaReturnsSingleton(t *testing.T) {
	first := NewCaptcha()
	second := NewCaptcha()

	if first == nil {
		t.Fatal("NewCaptcha() should return an initialized instance")
	}

	if first.captcha == nil {
		t.Fatal("NewCaptcha() should initialize the underlying captcha")
	}

	if first != second {
		t.Fatal("NewCaptcha() should return the same instance")
	}
}

func TestNewCaptchaReturnsSingletonConcurrently(t *testing.T) {
	const callers = 100

	captchas := make(chan *Captcha, callers)
	var waitGroup sync.WaitGroup
	waitGroup.Add(callers)

	for i := 0; i < callers; i++ {
		go func() {
			defer waitGroup.Done()
			captchas <- NewCaptcha()
		}()
	}

	waitGroup.Wait()
	close(captchas)

	first := NewCaptcha()
	for instance := range captchas {
		if instance != first {
			t.Fatal("NewCaptcha() should return the same instance to every goroutine")
		}
	}
}
