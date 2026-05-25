import React from 'react'
import LibraryScanButton from './LibraryScanButton'
import PruneMissingButton from './PruneMissingButton'

const LibraryListBulkActions = (props) => (
  <>
    <LibraryScanButton fullScan={false} {...props} />
    <LibraryScanButton fullScan={true} {...props} />
    <PruneMissingButton {...props} />
  </>
)

export default LibraryListBulkActions
