<script lang="ts">
  import { onMount } from 'svelte';
  import toast, { Toaster } from 'svelte-french-toast';

  let todos: Todo[] = [];
  let editingTodo: Todo | null;

  type Todo = {
    ID: number;
    Text: string;
    Completed: boolean;
  }

  onMount(async () => {
    const res = await fetch('http://localhost:8080/api/todos');
    todos = await res.json();
  });

  async function handleAddTodo(event: Event) {
    const input = (event.target as HTMLFormElement).querySelector('input[name="text"]') as HTMLInputElement;
    if (!input.value) return;
    const formData = new FormData(event.target as HTMLFormElement);
    await fetch('http://localhost:8080/api/todos', {
      method: 'POST',
      body: JSON.stringify({ text: formData.get('text') }),
      headers: { 'Content-Type': 'application/json' },
    });
    const res = await fetch('http://localhost:8080/api/todos');
    todos = await res.json();
    (event.target as HTMLFormElement).reset();
    toast.success("Todo added successfully");
  }

  function handleEditTodo(todo: Todo) {
    editingTodo = {...todo};
  }

   async function handleTodoComplete(todo) {
    const updatedTodo = {...todo, Completed: !todo.Completed};
    await fetch(`http://localhost:8080/api/todos/${todo.ID}`, {
      method: 'PUT',
      body: JSON.stringify(updatedTodo),
      headers: { 'Content-Type': 'application/json' },
    });
    const res = await fetch('http://localhost:8080/api/todos');
    todos = await res.json();
  }

  async function handleUpdateTodo(event: Event) {
    const formData = new FormData(event.target as HTMLFormElement);
    const updatedTodo = {
      ID: editingTodo.ID,
      Text: formData.get('text'),
      Completed: editingTodo.Completed
    };
    await fetch(`http://localhost:8080/api/todos/${editingTodo.ID}`, {
      method: 'PUT',
      body: JSON.stringify(updatedTodo),
      headers: { 'Content-Type': 'application/json' },
    });
    editingTodo = null;
    const res = await fetch('http://localhost:8080/api/todos');
    todos = await res.json();
    toast.success("Todo updated successfully");
  }

  async function handleDeleteTodo(id) {
    await fetch(`http://localhost:8080/api/todos/${id}`, { method: 'DELETE' });
    editingTodo = null;
    const res = await fetch('http://localhost:8080/api/todos');
    todos = await res.json();
    toast.success("Todo deleted successfully");
  }
</script>

<form on:submit|preventDefault={handleAddTodo}>
  <label for="add-todo">Add a todo:</label>
  <input id="add-todo" type="text" name="text" placeholder="Add a todo" />
  <button type="submit">Add</button>
</form>

{#if editingTodo}
  <form on:submit|preventDefault={handleUpdateTodo}>
    <label for="edit-todo">Edit a todo:</label>
    <input id="edit-todo" type="text" name="text" bind:value={editingTodo.Text} placeholder="Edit a todo" />
    <button type="submit">Update</button>
  </form>
{/if}

<ul>
  {#each todos as todo}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <li on:click={() => handleTodoComplete(todo)}>
      <span class:completed={todo.Completed}>{todo.Text}</span>
      <button on:click={() => handleEditTodo(todo)}>Edit</button>
      <button on:click={() => handleDeleteTodo(todo.ID)}>Delete</button>
    </li>
  {/each}
</ul>

<Toaster />

<style>
  ul {
    list-style: none;
    padding: 0;
  }

  li {
    cursor: pointer;
    border-radius: 4px;
  }

  li:hover {
    background-color: #333;
  }

  label[for="add-todo"], label[for="edit-todo"] {
    font-size: 1.2em;
    font-weight: bold;
  }

  input[type="text"] {
    padding: 0.5em;
    margin-bottom: 0.5em;
    border: 1px solid gray;
    border-radius: 4px;
  }

  .completed {
    text-decoration: line-through;
  }
</style>