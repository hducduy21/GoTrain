import htmx from 'htmx.org'
import Alpine from 'alpinejs'
import '@fortawesome/fontawesome-free/css/all.css';
import tasks from './components/tasks';

window.htmx = htmx
window.Alpine = Alpine
Alpine.start()

Alpine.data('tasks', tasks)