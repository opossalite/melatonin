import { useState } from "react";



interface Props {
    items: string[];
    heading: string;
}


function ListGroup({items, heading}: Props) {
    // hook
    const [selectedIndex, setSelectedIndex] = useState(-1);

    //usually for the key, item should have a .id
    return (
        <>
          <h1>{heading}</h1>
          {items.length === 0 && <p>No item found</p>}
          <ul className="list-group">
            {items.map((item, i) =>
                <li
                  className={i === selectedIndex ? 'list-group-item active' : 'list-group-item'}
                  key={item}
                  onClick={() => {setSelectedIndex(i)}}
                >
                  {item}
                </li>
            )}
          </ul>
        </>
    );
}

export default ListGroup;



