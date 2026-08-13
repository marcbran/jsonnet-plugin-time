local p = import 'pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/jsonnet-plugin-time',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'plugin/time',
  path: 'plugin/time',
  target: 'time',
}, |||
  Relative time resolution: get the current time, and add a signed duration spec to it. Calendar units - `Y`(ear), `M`(onth), `W`(eek), `D`(ay), uppercase - use calendar-aware arithmetic, so they land on the right day even though months and years aren't a fixed length. The remainder - `h`, `m`, `s`, `ms`, `us`, `ns`, lowercase - is fixed-length duration. Also parses absolute RFC3339 timestamps.
|||, {
  now: p.desc(|||
    Returns the current time as epoch milliseconds.
  |||),
  addDuration: p.desc(|||
    Adds a signed duration spec (e.g. `2h30m`, `1Y2M3D`, `-6h`) to `epochMs`, returning the resulting epoch milliseconds.
  |||),
  parseRFC3339: p.desc(|||
    Parses an RFC3339 timestamp (e.g. `2026-08-04T00:00:00Z`) into epoch milliseconds.
  |||),
})
