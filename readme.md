# Conversor de Temperaturas em Go

![Go Version](https://img.shields.io/badge/go-1.16+-blue)
![License](https://img.shields.io/badge/license-MIT-green)

Um conversor de temperaturas CLI eficiente escrito em Go que realiza conversões precisas entre Celsius, Fahrenheit e Kelvin.

## 📦 Instalação

### Via go install
```bash
go install github.com/seu-usuario/conversor-temperatura@latest
Compilação manual
bash
git clone https://github.com/seu-usuario/conversor-temperatura.git
cd conversor-temperatura
go build -o conversor-temp
```
🚀 Como Usar
Sintaxe básica
```bash
./conversor-temp -valor <temperatura> -de <unidade_origem> -para <unidade_destino>
```
Exemplos
```bash
# Converter 100°C para Fahrenheit
./conversor-temp -valor 100 -de C -para F

# Converter 32°F para Kelvin
./conversor-temp -valor 32 -de F -para K

# Converter 273.15K para Celsius
./conversor-temp -valor 273.15 -de K -para C
```
Opções

| Flag      | Descrição                          | Valores Aceitos | Padrão |
|-----------|------------------------------------|-----------------|--------|
| `-valor`  | Valor da temperatura a converter   | Número real     | 0      |
| `-de`     | Unidade de origem                  | C, F, K         | C      |
| `-para`   | Unidade de destino                 | C, F, K         | F      |
| `-h`      | Mostrar ajuda                      | -               | -      |

📊 Conversões Suportadas
| De/Para   | Celsius (°C) | Fahrenheit (°F) | Kelvin (K) |
|-----------|--------------|-----------------|------------|
| **Celsius** | -            | ✅               | ✅          |
| **Fahrenheit** | ✅          | -               | ✅          |
| **Kelvin** | ✅           | ✅               | -          |

Build
```bash
go build -o conversor-temp
```
