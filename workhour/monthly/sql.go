package monthly

const sqlFindMonthly = `
  SELECT 
    id, 
    user_id, 
    TO_CHAR(work_day, 'YYYY-MM-DD'), 
    TO_CHAR(hello, 'HH24:MI'), 
    TO_CHAR(goodbye, 'HH24:MI'), 
    TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'), 
    TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
  FROM workhour
  WHERE work_day >= TO_DATE($1, 'YYYY-MM-DD') 
    AND work_day <= TO_DATE($2, 'YYYY-MM-DD') 
    AND user_id = $3 
  ORDER BY work_day`
