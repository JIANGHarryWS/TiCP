/*!
 * Copyright (C) 2016-present, Yuansuan.cn
 */

import React from 'react'
import { Tag as AntTag, Divider as AntDivider } from 'antd'
import Table from '..'

const data = [
  {
    key: '1',
    name: 'John Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer'],
  },
  {
    key: '2',
    name: 'Jim Green',
    age: 42,
    address: 'London No. 1 Lake Park',
    tags: ['loser'],
  },
  {
    key: '3',
    name: 'Joe Black',
    age: 32,
    address: 'Sidney No. 1 Lake Park',
    tags: ['cool', 'teacher'],
  },
]

export const Customizable = () => (
  <Table
    tableId='custom_table'
    defaultConfig={[
      { key: 'name', active: false },
      { key: 'age', active: true },
      { key: 'address', active: true },
      { key: 'tags', active: true },
      { key: 'action', active: true },
    ]}
    props={{ autoHeight: true, data }}
    columns={[
      {
        header: 'Name',
        dataKey: 'name',
        props: {
          flexGrow: 1,
        },
      },
      {
        header: 'Age',
        dataKey: 'age',
        props: {
          flexGrow: 1,
        },
      },
      {
        header: 'Address',
        dataKey: 'address',
        props: {
          flexGrow: 1,
        },
      },
      {
        header: 'Tags',
        props: {
          flexGrow: 1,
        },
        cell: {
          props: {
            dataKey: 'tags',
          },
          render: ({ rowData, dataKey }) => {
            const tags = rowData[dataKey]

            return (
              <span>
                {tags.map(tag => {
                  let color = tag.length > 5 ? 'geekblue' : 'green'
                  if (tag === 'loser') {
                    color = 'volcano'
                  }
                  return (
                    <AntTag color={color} key={tag}>
                      {tag.toUpperCase()}
                    </AntTag>
                  )
                })}
              </span>
            )
          },
        },
      },
      {
        header: 'Action',
        props: {
          flexGrow: 1,
        },
        cell: {
          props: {
            dataKey: 'action',
          },
          render: ({ rowData }) => {
            return (
              <span>
                <a>Invite {rowData.name}</a>
                <AntDivider type='vertical' />
                <a>Delete</a>
              </span>
            )
          },
        },
      },
    ]}
  />
)
