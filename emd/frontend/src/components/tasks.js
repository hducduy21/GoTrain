export default () => ({
    tasks: [],
    newTaskName: '',
 
    init(){
        fetch('/api/tasks')
            .then(response => response.json())
            .then(data => {
                console.log(data)
                this.tasks = data
            })
    },

    createTask(){
        fetch('/api/tasks', {method: 'POST', body: JSON.stringify({name: this.newTaskName})})
        .then(response => response.json()).then(data =>{
            this.tasks.push(data)
            this.newTaskName = ''
        })
    },

    updateTask(id, name){
        fetch(`/api/tasks/${id}`, {method: 'POST', body: JSON.stringify({name})})
        .then(response => response.json()).then(data => {
            const index = this.tasks.findIndex(t => t.id === id)
            this.tasks[index] = data
        })
    },

    deleteTask(id){
        fetch(`/api/tasks/${id}`, {method: 'DELETE'})
        .then(response => {
            if(response.ok){
                this.tasks = this.tasks.filter(t => t.id !== id)
            }
        })
    },

    joinTask(id){
        fetch(`/api/tasks/${id}/join`, {method: 'PATCH'})
        .then(response => response.json()).then(data =>{
            const index = this.tasks.findIndex(t => t.id === id)
            if(this.tasks[index].edges.users === undefined){
                this.tasks[index].edges.users = [data]
            }else{
                this.tasks[index].edges.users.push(data)
            }
        })
    },

    leaveTask(id){
        fetch(`/api/tasks/${id}/leave`, {method: 'PATCH'})
        .then(response => response.json()).then(data =>{
            const index = this.tasks.findIndex(t => t.id === id)
            if(this.tasks[index].edges.users === undefined){
                this.tasks[index].edges.users = [data]
            }else{
                this.tasks[index].edges.users = this.tasks[index].edges.users.filter(u => u.id !== data.id)
            }
        })
    },

    updateTaskStatus(id, status){
        fetch(`/api/tasks/${id}/status`, {method: 'PATCH', body: JSON.stringify({done: status})})
        .then(response => response.json()).then(data =>{
            console.log(data)
        })
    }
})