-- Cria a tabela de feature flags no banco de dados 'mydb'
CREATE TABLE IF NOT EXISTS feature_flags (
                                             flag_name VARCHAR(50) PRIMARY KEY,
                                             enabled BOOLEAN NOT NULL
);

-- Exemplo de inserção de algumas feature flags iniciais
INSERT INTO feature_flags (flag_name, enabled) VALUES
                                                   ('processOrderFeature', true),
                                                   ('featureB', false),
                                                   ('featureC', true)
ON CONFLICT (flag_name) DO NOTHING;
