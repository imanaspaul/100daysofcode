import React , { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { addtodo , allTodos, removetodo} from './todoSlice';

export function Todo(){
    const todos = useSelector(allTodos)
    const dispatch = useDispatch();
    return(
        <>
        <h1>Todo add</h1>
        <pre>{todos}</pre>
         <button
          aria-label="Add todo"
          onClick={() => dispatch(addtodo(1))}
        >
          +
        </button>
         <button
          aria-label="remove todo"
          onClick={() => dispatch(removetodo())}
        >
          remove todos
        </button>
        </>
    )
}