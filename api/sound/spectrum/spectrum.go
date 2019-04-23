package spectrum

import (
	"errors"
	"io"
	"math"
	"math/cmplx"

	"github.com/go-audio/wav"
)

// FrequencyUpperBound is the maximum frequency analyzed. All frequencies beyond this will be discarded
const FrequencyUpperBound = 47000

// freqStep is the delta frequency step
const freqStep = 1

// FrequencyPower is the representation of a frequency power
type FrequencyPower struct {
	Freq  int
	Value float64
}

// FrequenciesAnalysis modes an analysis of frequencies for an audio file
type FrequenciesAnalysis struct {
	Spectrum []float64
	MaxFreq  int
	MinFreq  int
	MaxPower FrequencyPower
}

// PCMFromWav reads the PCM as Buffer and the sample rate
func PCMFromWav(wavReader io.ReadSeeker) ([]int, int, error) {
	d := wav.NewDecoder(wavReader)
	// TODO: use Buffer and read samples into chunks rather than loading all this into memory
	pcm, err := d.FullPCMBuffer()
	if err != nil {
		return nil, 0, err
	}
	return pcm.Data, pcm.Format.SampleRate, nil
}

// // NormalizedFrequenciesSpectrum calculates the normalized frequencies spectrum from a PCM sample buffer
// func NormalizedFrequenciesSpectrum(samples []int, sampleRate int) (FrequenciesAnalysis, error) {
// 	freqAnalysis, err := FrequenciesSpectrum(samples, sampleRate)
// 	if err != nil {
// 		return freqAnalysis, err
// 	}
// 	freqAnalysis.Spectrum, err = normalizeFrequenciesSpectrum(freqAnalysis.Spectrum)
// 	return freqAnalysis, err
// }

// FrequenciesSpectrumAnalysis calculates the frequencies spectrum from a PCM sample buffer
func FrequenciesSpectrumAnalysis(samples []int, sampleRate int) (FrequenciesAnalysis, error) {
	freqAnalysis := FrequenciesAnalysis{}
	complexSamples := toComplexSamples(samples)
	err := fftTransform(&complexSamples)
	if err != nil {
		return freqAnalysis, err
	}
	freqSamplingStep := float64(sampleRate) / float64(len(complexSamples))
	scaleStep := float64(freqStep) / freqSamplingStep
	freqLimit := float64(sampleRate) / 2.0

	freqAnalysis.Spectrum = make([]float64, FrequencyUpperBound)
	freqAnalysis.MaxFreq = -1
	freqAnalysis.MinFreq = FrequencyUpperBound
	max := 0.0
	for i := 0.0; i*freqSamplingStep <= freqLimit; i += scaleStep {
		valueDiscarded := false
		abs := cmplx.Abs(complexSamples[int(i)])
		freq := int(i * freqSamplingStep)
		if freq <= FrequencyUpperBound {
			if abs < 0.0001 {
				valueDiscarded = true
				freqAnalysis.Spectrum[freq] = 0
			} else {
				freqAnalysis.Spectrum[freq] = abs
				if max < abs {
					max = abs
					freqAnalysis.MaxPower = FrequencyPower{Freq: freq, Value: abs}
				}
			}
		} else {
			valueDiscarded = true
		}
		if !valueDiscarded {
			if freq > freqAnalysis.MaxFreq {
				freqAnalysis.MaxFreq = freq
			}
			if freq < freqAnalysis.MinFreq {
				freqAnalysis.MinFreq = freq
			}
		}
	}

	return freqAnalysis, nil
}

func toComplexSamples(samples []int) []complex128 {
	complexSamples := make([]complex128, nearestBiggerPowerOf2(uint(len(samples))))
	for i, s := range samples {
		complexSamples[i] = complex(float64(s), 0)
	}
	return complexSamples
}

// func normalizeFrequenciesSpectrum(spectrum []float64, ) (map[int]float64, error) {
// 	maxValue := maxValue(spectrum)
// 	if maxValue == 0.0 {
// 		return nil, errors.New("can't normalize spectrum. Max value is 0")
// 	}
// 	for k, v := range spectrum {
// 		spectrum[k] = v / maxValue
// 	}
// 	return spectrum, nil
// }

// fftTransform is a fast fourier transform
func fftTransform(samples *[]complex128) error {
	n := len(*samples)
	if n <= 1 {
		return errors.New("fast fourier transform requires at least one sample")
	}
	a0 := make([]complex128, n/2)
	a1 := make([]complex128, n/2)
	for i, j := 0, 0; i < n; i, j = i+2, j+1 {
		a0[j] = (*samples)[i]
		a1[j] = (*samples)[i+1]
	}
	// ignoring errors int the following two lines
	_ = fftTransform(&a0)
	_ = fftTransform(&a1)
	angle := 2 * math.Pi / float64(n)
	w := complex128(1)
	wn := complex(math.Cos(angle), math.Sin(angle))
	for i := 0; i < n/2; i++ {
		(*samples)[i] = a0[i] + w*a1[i]
		(*samples)[i+n/2] = a0[i] - w*a1[i]
		w *= wn
	}
	return nil
}

func nearestBiggerPowerOf2(x uint) uint {
	x = x - 1
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	return x + 1
}
