import React, {useState} from 'react';
import './App.css';
import { useApiTask } from './hooks/useApiTask';

function App() {
  const api = useApiTask();
  const [tasks, setTasks] = useState<any[]>();

  const init = async () => {
    const sss =  await api.getTasks();
    console.log(sss);
    setTasks(sss);
  };

  React.useEffect(() => {
    init();
  }, []);


  return (
    <div className="App">
      <div>
      {tasks && tasks.map(
        (task, index) => <div key={index}>{task.id} {task.name}</div>
      )}
      </div>
    </div>
  );
}

export default App;
