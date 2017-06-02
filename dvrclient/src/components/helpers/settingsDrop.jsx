import React from 'react'
import { Dropdown } from 'semantic-ui-react'

const SettingsDrop = () => (
  <Dropdown className="headerico" icon='setting'>
    <Dropdown.Menu>
      <Dropdown.Item text='Sign-Out' />
    </Dropdown.Menu>
  </Dropdown>
)

export default SettingsDrop