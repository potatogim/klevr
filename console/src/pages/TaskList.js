import React, { useState } from 'react';
import { x } from '@xstyled/emotion';
import { Box, Container } from '@material-ui/core';
import { Button } from 'antd';

import AllTasks from 'src/components/task/AllTasks';
import OrderList from 'src/components/task/OrderList';
import SchedulerList from 'src/components/task/SchedulerList';
import ProvisioningList from 'src/components/task/ProvisioningList';

import AddTask from 'src/components/task/AddTask';
import Refresh from 'src/components/common/Refresh';
import TaskLog from 'src/components/task/TaskLog';

const content = [
  {
    tab: 'All',
    content: <AllTasks />
  },
  {
    tab: 'Order',
    content: <OrderList />
  },
  {
    tab: 'Scheduler',
    content: <SchedulerList />
  },
  {
    tab: 'Provisioning',
    content: <ProvisioningList />
  }
];

const useTabs = (initialTabs, allTabs) => {
  const [contentIndex, setContentIndex] = useState(initialTabs);
  return {
    contentItem: allTabs[contentIndex],
    contentChange: setContentIndex
  };
};

const TaskList = () => {
  const { contentItem, contentChange } = useTabs(0, content);

  return (
    <>
      <Box
        sx={{
          backgroundColor: 'background.default',
          minHeight: '100%',
          py: 3
        }}
      >
        <Container maxWidth={false}>
          <x.div
            display="flex"
            justifyContent="space-between"
            alignItems="center"
            mb="20"
          >
            <div>
              {content.map((section, index) => (
                <Button
                  onClick={() => contentChange(index)}
                  key={index}
                  type="text"
                >
                  <x.span
                    color={
                      section.tab === contentItem.tab
                        ? 'rgb(25, 118, 210)'
                        : 'black'
                    }
                    fontWeight="500"
                    lineHeight="1.75"
                    fontSize="0.875rem"
                  >
                    {section.tab.toUpperCase()}
                  </x.span>
                </Button>
              ))}
            </div>
            <x.div w="100px" display="flex" justifyContent="space-between">
              <AddTask />
              <Refresh from="task" />
            </x.div>
          </x.div>
          {contentItem.content}
        </Container>
      </Box>
    </>
  );
};

export default TaskList;
