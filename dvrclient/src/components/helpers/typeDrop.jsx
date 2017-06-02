import React from 'react'
import { Dropdown } from 'semantic-ui-react'


const typeOptions = [
{ key: 'sf', value: 'Semi-Formal', text: 'Semi-Formal' },
{ key: 'cas', value: 'Casual', text: 'Casual' },
{ key: 'fes', value: 'Festive', text: 'Festive' },
{ key: 'bt', value: 'Black Tie', text: 'Black Tie' },
{ key: 'wt', value: 'White Tie', text: 'White Tie' },
{ key: 'other', value: 'Other', text: 'Other' }
]

const TypeDown = () => (
  <Dropdown id="typeDown" fluid closeOnChange placeholder='Type' compact selection options={typeOptions} />
)

export default TypeDown;