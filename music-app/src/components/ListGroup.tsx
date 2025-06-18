import {MouseEvent} from "react";



function ListGroup() {
    let items = [
        'New York',
        'San Francisco',
        'Tokyo',
        'London',
        'Paris',
    ];
    let selectedIndex = -1;

    // event handler
    const handleClick = (event: MouseEvent) => console.log(event);
    

    //usually for the key, item should have a .id
    return (
        <>
          <h1>List</h1>
          {items.length === 0 && <p>No item found</p>}
          <ul className="list-group">
            {items.map((item, i) =>
                <li
                  className={i === selectedIndex ? 'list-group-item' : 'active'}
                  key={item}
                  onClick={handleClick}
                >
                  {item}
                </li>
            )}
          </ul>
        </>
    );
}

export default ListGroup;



