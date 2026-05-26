import PropTypes from 'prop-types'
import React from 'react'
import { usePermissions, useRecordContext } from 'react-admin'
import config from '../config'

export const formatPath = (record, permissions) => {
  if (!record?.path) {
    return ''
  }

  let path = permissions === 'admin' ? record.libraryPath : ''

  if (path && path.endsWith(config.separator)) {
    return `${path}${record.path}`
  }
  return path ? `${path}${config.separator}${record.path}` : record.path
}

export const PathField = (props) => {
  const record = useRecordContext(props)
  const { permissions } = usePermissions()
  const path = formatPath(record, permissions)

  return <span>{path}</span>
}

PathField.propTypes = {
  record: PropTypes.object,
}
