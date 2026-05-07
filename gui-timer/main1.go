package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type TimerApp struct {
	minutesEntry   *widget.Entry
	secondsEntry   *widget.Entry
	startButton    *widget.Button
	stopButton     *widget.Button
	timerLabel     *widget.Label
	isRunning      bool
	ctx            context.Context
	cancel         context.CancelFunc
	totalSeconds   int
	currentSeconds int
}

// Создаем собственный генератор синусоиды
type sineWave struct {
	freq   float64
	sample float64
	sr     beep.SampleRate
}

func newSineWave(sr beep.SampleRate, freq float64) *sineWave {
	return &sineWave{
		freq: freq,
		sr:   sr,
	}
}

func (s *sineWave) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		val := math.Sin(s.sample * 2 * math.Pi * s.freq / float64(s.sr))
		samples[i][0] = val
		samples[i][1] = val
		s.sample++
	}
	return len(samples), true
}

func (s *sineWave) Err() error {
	return nil
}

func NewTimerApp() *TimerApp {
	app := &TimerApp{}
	app.setupUI()
	app.initAudio()
	return app
}

func (t *TimerApp) initAudio() {
	// Инициализация аудио системы
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
}

func (t *TimerApp) setupUI() {
	// Создаем поля ввода
	t.minutesEntry = widget.NewEntry()
	t.minutesEntry.SetPlaceHolder("Минуты")
	t.minutesEntry.SetText("1")

	t.secondsEntry = widget.NewEntry()
	t.secondsEntry.SetPlaceHolder("Секунды")
	t.secondsEntry.SetText("30")

	// Создаем кнопки
	t.startButton = widget.NewButton("Старт", t.startTimer)
	t.stopButton = widget.NewButton("Стоп", t.stopTimer)
	t.stopButton.Disable()

	// Создаем лейбл для отображения таймера
	t.timerLabel = widget.NewLabel("00:00")
	t.timerLabel.Alignment = fyne.TextAlignCenter
	t.timerLabel.TextStyle = fyne.TextStyle{Bold: true}
}

func (t *TimerApp) startTimer() {
	if t.isRunning {
		return
	}

	// Получаем значения из полей ввода
	minutesStr := t.minutesEntry.Text
	secondsStr := t.secondsEntry.Text

	minutes, err1 := strconv.Atoi(minutesStr)
	seconds, err2 := strconv.Atoi(secondsStr)

	if err1 != nil || err2 != nil || minutes < 0 || seconds < 0 {
		t.updateTimerLabel("Ошибка ввода!")
		return
	}

	t.totalSeconds = minutes*60 + seconds
	if t.totalSeconds == 0 {
		t.updateTimerLabel("Время должно быть > 0")
		return
	}

	t.isRunning = true
	t.ctx, t.cancel = context.WithCancel(context.Background())
	
	// Обновляем UI в главном потоке
	t.startButton.Disable()
	t.stopButton.Enable()
	t.minutesEntry.Disable()
	t.secondsEntry.Disable()

	// Запускаем таймер в отдельной горутине
	go t.runTimer()
}

func (t *TimerApp) stopTimer() {
	if !t.isRunning {
		return
	}

	t.isRunning = false
	if t.cancel != nil {
		t.cancel()
	}

	// Обновляем UI в главном потоке
	t.startButton.Enable()
	t.stopButton.Disable()
	t.minutesEntry.Enable()
	t.secondsEntry.Enable()
	t.updateTimerLabel("Остановлено")
}

func (t *TimerApp) updateTimerLabel(text string) {
	// Используем fyne.Do для безопасного обновления UI из горутины
	fyne.Do(func() {
		t.timerLabel.SetText(text)
	})
}

func (t *TimerApp) runTimer() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for t.isRunning {
		t.currentSeconds = t.totalSeconds

		// Обратный отсчет
		for t.currentSeconds > 0 && t.isRunning {
			select {
			case <-t.ctx.Done():
				return
			case <-ticker.C:
				// Обновляем отображение времени в главном потоке
				minutes := t.currentSeconds / 60
				seconds := t.currentSeconds % 60
				timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)
				t.updateTimerLabel(timeStr)
				t.currentSeconds--
			}
		}

		if t.isRunning && t.currentSeconds <= 0 {
			// Время истекло - подаем звуковой сигнал
			t.playBeep()
			t.updateTimerLabel("ВРЕМЯ!")
			
			// Небольшая пауза перед перезапуском
			select {
			case <-t.ctx.Done():
				return
			case <-time.After(1 * time.Second):
				// Продолжаем цикл для автоматического перезапуска
			}
		}
	}
}

func (t *TimerApp) playBeep() {
	// Создаем короткий звуковой сигнал (гудок)
	sr := beep.SampleRate(44100)
	
	// Создаем синусоидальную волну частотой 800 Гц
	sine := newSineWave(sr, 800)
	
	// Ограничиваем длительность звука до 0.3 секунды
	beepSound := beep.Take(sr.N(time.Millisecond*300), sine)
	
	// Воспроизводим звук
	speaker.Play(beepSound)
}

func (t *TimerApp) getContainer() *fyne.Container {
	// Создаем форму с полями ввода
	inputForm := container.NewGridWithColumns(2,
		widget.NewLabel("Минуты:"), t.minutesEntry,
		widget.NewLabel("Секунды:"), t.secondsEntry,
	)

	// Создаем контейнер с кнопками
	buttonContainer := container.NewGridWithColumns(2,
		t.startButton, t.stopButton,
	)

	// Основной контейнер
	mainContainer := container.NewVBox(
		widget.NewCard("Настройки таймера", "", inputForm),
		widget.NewSeparator(),
		widget.NewCard("", "", container.NewCenter(t.timerLabel)),
		widget.NewSeparator(),
		buttonContainer,
	)

	return mainContainer
}

func main() {
	// Создаем приложение
	myApp := app.New()
	myApp.SetIcon(nil)

	// Создаем окно
	myWindow := myApp.NewWindow("Таймер")
	myWindow.Resize(fyne.NewSize(300, 280))
	myWindow.CenterOnScreen()

	// Создаем экземпляр нашего приложения
	timerApp := NewTimerApp()

	// Устанавливаем содержимое окна
	myWindow.SetContent(timerApp.getContainer())

	// Обработчик закрытия окна
	myWindow.SetCloseIntercept(func() {
		if timerApp.isRunning {
			timerApp.stopTimer()
		}
		myWindow.Close()
	})

	// Показываем окно и запускаем приложение
	myWindow.ShowAndRun()
}

