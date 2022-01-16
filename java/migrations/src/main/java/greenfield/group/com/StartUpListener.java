package greenfield.group.com;

import io.quarkus.runtime.ApplicationLifecycleManager;
import io.quarkus.runtime.Quarkus;
import io.quarkus.runtime.StartupEvent;

import javax.enterprise.context.ApplicationScoped;
import javax.enterprise.event.Observes;
import javax.inject.Inject;

/**
 * Бин прослушивания поднятия контекста
 *
 */
@ApplicationScoped
public class StartUpListener {

    void onStart(@Observes StartupEvent ev) {
        ApplicationLifecycleManager.exit(0);
    }

}
