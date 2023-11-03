-- Create a database
CREATE DATABASE mydatabase;

-- Connect to the newly created database
\c mydatabase

-- Create a table to store news
CREATE TABLE news (
    title TEXT,
    description TEXT,
    timestamp BIGINT,
    id TEXT
);

-- Insert news data into the table
INSERT INTO news (title, description, timestamp, id)
VALUES
    ('World Leaders Gather for Climate Summit', 'Leaders from around the world converged in Glasgow for a crucial climate summit, with discussions focused on addressing the global climate crisis.', 1667251200, '1'),
    ('Tech Giant Apple Unveils New iPhone 15', 'Apple Inc. announced the release of its latest flagship smartphone, the iPhone 15, boasting advanced features and improved performance.', 1667254200, '2'),
    ('Scientists Discover New Species in Amazon Rainforest', 'Researchers exploring the Amazon rainforest have identified a previously unknown species of amphibian, adding to the regions incredible biodiversity.', 1667257200, '3'),
    ('Stock Markets Experience Volatility Amid Economic Uncertainty', 'Global stock markets saw significant fluctuations as concerns about inflation and economic instability continue to weigh on investor sentiment.', 1667260200, '4'),
    ('Historic Peace Agreement Signed in Conflict-Stricken Region', 'A long-awaited peace accord was signed in a conflict-stricken region, bringing hope for an end to years of violence and instability.', 1667263200, '5');