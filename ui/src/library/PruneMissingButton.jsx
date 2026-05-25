import React, { useState } from 'react'
import PropTypes from 'prop-types'
import { Button, useNotify, useRefresh, useTranslate } from 'react-admin'
import DeleteSweepIcon from '@material-ui/icons/DeleteSweep'
import { httpClient } from '../dataProvider'
import { REST_URL } from '../consts'

const PruneMissingButton = ({ className }) => {
  const [loading, setLoading] = useState(false)
  const notify = useNotify()
  const refresh = useRefresh()
  const translate = useTranslate()

  const handleClick = async () => {
    setLoading(true)
    try {
      await httpClient(`${REST_URL}/missing/prune`, { method: 'POST' })
      notify('resources.library.notifications.pruneMissingCompleted', 'info')
      refresh()
    } catch (error) {
      notify('resources.library.notifications.pruneMissingError', 'warning')
    } finally {
      setLoading(false)
    }
  }

  return (
    <Button
      onClick={handleClick}
      disabled={loading}
      label={translate('resources.library.actions.pruneMissing')}
      className={className}
    >
      <DeleteSweepIcon />
    </Button>
  )
}

PruneMissingButton.propTypes = {
  className: PropTypes.string,
}

export default PruneMissingButton
