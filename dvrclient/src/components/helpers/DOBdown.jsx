import React from 'react'
import { Form, Input } from 'semantic-ui-react'

const DOBForm = () => (
    <Form.Group inline>
        <Form.Field>
            <Input id='dobDay' type='number' maxLength='2' placeholder='Day' />
        </Form.Field>
        <Form.Field>
            <Input id='dobYear' type='number' maxLength='4' placeholder='Year' />
        </Form.Field>
    </Form.Group>
)

export default DOBForm
