import React from 'react';

function PersistentState(key: string, defaultValue: any) {
  const storedValue = localStorage.getItem(key);
  const [value, setValue] = React.useState(
    (storedValue && JSON.parse(storedValue)) || defaultValue
  );

  React.useEffect(() => {
    localStorage.setItem('isAuth', JSON.stringify(value));
  }, [value]);

  return [value, setValue];
};

export default PersistentState;
