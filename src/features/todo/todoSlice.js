import { createSlice } from '@reduxjs/toolkit';

export const todoSlice = createSlice({
    name: 'todo',
    initialState : {
        todos : [1,3,4]
    },
    reducers : {
        addtodo : (state, action) => {
            state.todos.push(action.payload);
        },
        removetodo : state=>  {
            state.todos = [1]
        }
    }
});

export const { addtodo, removetodo } = todoSlice.actions
export const allTodos = state => state.todo.todos
export default todoSlice.reducer;