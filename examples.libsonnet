local p = import 'pkg/main.libsonnet';

p.ex({
  addDuration: p.ex([{
    name: 'compound fixed-length duration',
    inputs: [0, '2h30m'],
  }, {
    name: 'calendar-aware month, lands on the right day',
    inputs: [0, '1M'],
  }, {
    name: 'negative duration',
    inputs: [0, '-90s'],
  }]),
})
