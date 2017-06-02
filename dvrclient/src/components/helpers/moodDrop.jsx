import React from 'react'
import { Dropdown } from 'semantic-ui-react'

const moodOptions = [
{ key: 'fun', value: 'Funny', text: 'Fun' },
{ key: 'silly', value: 'Silly', text: 'Silly' },
{ key: 'fancy', value: 'Fancy', text: 'Fancy' },
{ key: 'relaxed', value: 'Relaxed', text: 'Relaxed' },
{ key: 'focused', value: 'Focused', text: 'Focused' },
]

const MoodDown = () => (
  <Dropdown id="moodDown" fluid closeOnChange placeholder='Mood' compact selection options={moodOptions} />
)

export default MoodDown;