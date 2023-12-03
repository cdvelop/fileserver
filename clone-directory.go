package fileserver

import (
	"fmt"
	"os/exec"
	"runtime"
)

// robocopy permite hacer copias espejo de directorios,
// es decir sincronizar dos directorios.
// Esto es muy útil para hacer copias de respaldo,
// ya que tras la copia, robocopy eliminará del directorio de destino
// los archivos que ya no existan en el directorio de origen.
// Por ejemplo, para copiar
// C:\Usuarios\Mis documentos a D:\backup\Mis documentos,
// ejecuta el comando robocopy “C:\Usuarios\Mis documentos” “D:\backup\Mis documentos” /mir /z.
// El modificador /mir es el que permite el modo espejo.
// Por su parte el modificador /z permitirá reanudar la copia en caso de que se produzca una interrupción,
// ya sea por corte de energía u otro motivo.

func CloneDirectory(origin, destiny string) (err string) {

	switch runtime.GOOS {
	case "windows":
		return cloneWindowsFolder(origin, destiny)
	default:
		return "BACKUP ARCHIVOS LINUX NO IMPLEMENTADO"
	}

}

func cloneWindowsFolder(origin, destiny string) (err string) {
	var change = "sin cambios"
	_, er := exec.Command("robocopy", origin, destiny, "/mir", "/z").Output()
	if er != nil {
		change = "archivos nuevos " + er.Error()
	}

	fmt.Printf(">>> RESPALDO DIRECTORIO: %v A: %v [%v]\n", origin, destiny, change)
	return
}
