local p = import 'pkg/main.libsonnet';

p.ex({
  now: p.ex([{
    name: 'current time',
    inputs: [],
  }]),
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
  parse: p.ex([{
    name: 'rfc3339 utc timestamp',
    inputs: ['2026-08-04T00:00:00Z', '2006-01-02T15:04:05Z07:00'],
  }, {
    name: 'timestamp with an explicit offset',
    inputs: ['2026-08-04T02:00:00+02:00', '2006-01-02T15:04:05Z07:00'],
  }]),
  format: p.ex([{
    name: 'date and time',
    inputs: [1754265600000, '2006-01-02 15:04'],
  }]),
})
